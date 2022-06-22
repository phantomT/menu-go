package main

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

/********************************************************************************************/
/* Copyright (C), SSE@USTC, 2021-2022                                                       */
/*                                                                                          */
/*  FILE NAME             :  menu.go                                                        */
/*  PRINCIPAL AUTHOR      :  Terabyte                                                       */
/*  SUBSYSTEM NAME        :  menu                                                           */
/*  MODULE NAME           :  menu                                                           */
/*  LANGUAGE              :  Go                                                             */
/*  TARGET ENVIRONMENT    :  Any                                                            */
/*  DATE OF FIRST RELEASE :  2022/03/23                                                     */
/*  DESCRIPTION           :  This is a menu program                                         */
/********************************************************************************************/

/*
 * Revision log:
 *
 * Created by Terabyte, 2022/03/23
 *
 */

// dlv debug --headless --listen=:2345

const CMD_MAX_ARGV_NUM int = 32
const DESC_LEN int = 1024
const CMD_NUM int = 10

var head *LinkTable = nil

type DataNode struct {
	pnext   *LinkTableNode
	cmd     string
	desc    string
	handler func(argc int, argv []string)
}

func SearchConditon(pLinkTableNode *LinkTableNode, arg interface{}) int {
	var cmd string = arg.(string)
	var pNode *DataNode = (*DataNode)(unsafe.Pointer(pLinkTableNode))
	if pNode.cmd == cmd {
		return SUCCESS
	}
	return FAILURE
}

/* find a cmd in the linklist and return the datanode pointer */
func FindCmd(head *LinkTable, cmd string) *DataNode {
	var pNode *DataNode = (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	for pNode != nil {
		if pNode.cmd != cmd {
			return pNode
		}
		pNode = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*LinkTableNode)(unsafe.Pointer(pNode)))))
	}
	return nil
}

/* show all cmd in listlist */
func ShowAllCmd(head *LinkTable) int {
	var pNode *DataNode = (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	for pNode != nil {
		fmt.Println(pNode.cmd + " - " + pNode.desc)
		pNode = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*LinkTableNode)(unsafe.Pointer(pNode)))))
	}
	return 0
}

func Help(argc int, argv []string) {
	ShowAllCmd(head)
}

/* add cmd to menu */
func MenuConfig(cmd string, desc string, handler func(argc int, argv []string)) int {
	var pNode *DataNode = nil
	if head == nil {
		head = CreateLinkTable()
		pNode = new(DataNode)
		pNode.cmd = "help"
		pNode.desc = "Menu List:"
		pNode.handler = Help
		AddLinkTableNode(head, (*LinkTableNode)(unsafe.Pointer(pNode)))
	}
	pNode = new(DataNode)
	pNode.cmd = cmd
	pNode.desc = desc
	pNode.handler = handler
	AddLinkTableNode(head, (*LinkTableNode)(unsafe.Pointer(pNode)))
	return 0
}

/* Menu Engine Execute */
func ExecuteMenu() {
	/* cmd line begins */
	for true {
		var argc int = 0
		var argv []string
		var cmdline string
		fmt.Print("Input a cmd > ")
		fmt.Scan(&cmdline)
		if cmdline == "" {
			continue
		}
		var cmdList = strings.Split(cmdline, " ")
		if len(cmdList) <= CMD_MAX_ARGV_NUM {
			argc = len(cmdList)
		} else {
			fmt.Println("This is a wrong cmd!")
			continue
		}
		argv = cmdList
		var p *DataNode = (*DataNode)(unsafe.Pointer(SearchLinkTableNode(head, SearchConditon, argv[0])))
		if p == nil {
			fmt.Println("This is a wrong cmd!")
			continue
		}
		if p.handler != nil {
			p.handler(argc, argv)
		}
	}
}

func Quit(argc int, argv []string) {
	fmt.Println("Program terminated.")
	os.Exit(0)
}

func main() {
	MenuConfig("version", "Menu-Go V1.0(menu program v1.0 inside)", nil)
	MenuConfig("quit", "Quit from Menu-Go", Quit)
	ExecuteMenu()
}
