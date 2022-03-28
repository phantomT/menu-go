package main

/********************************************************************************************/
/* Copyright (C), SSE@USTC, 2021-2022                                                  		*/
/*                                                                                          */
/*  FILE NAME             :  menu.go                                                        */
/*  PRINCIPAL AUTHOR      :  Terabyte                                                       */
/*  SUBSYSTEM NAME        :  menu                                                           */
/*  MODULE NAME           :  menu                                                           */
/*  LANGUAGE              :  Go                                                             */
/*  TARGET ENVIRONMENT    :  Any                                                     		*/
/*  DATE OF FIRST RELEASE :  2022/03/23                                                     */
/*  DESCRIPTION           :  This is a menu program                                         */
/********************************************************************************************/

/*
 * Revision log:
 *
 * Created by Terabyte, 2022/03/23
 *
 */

import (
	"fmt"
	"os"
)

// dlv debug --headless --listen=:2345

const CMD_MAX_LEN int = 128
const DESC_LEN int = 1024
const CMD_NUM int = 10

type DataNode struct {
	cmd     string
	desc    string
	handler func()
	next    *DataNode
}

var head [5]DataNode

func init() {
	head = [5]DataNode{
		{"help", "Get the Menu List", Help, &head[1]},
		{"list", "Hurry up to make your choise!", nil, &head[2]},
		{"setup", "Watch up where you going", nil, &head[3]},
		{"quit", "Exit Program", Quit, &head[4]},
		{"nut", "Easter Egg", Nut, nil},
	}
}

func main() {
	var cmdline string
	fmt.Println("WelCome to The Machine")
	fmt.Println("####################################")
	for true {
		fmt.Println("\n#*********# Go_MENU v0.1 #*********#")
		fmt.Println("Input a cmd > ")
		fmt.Scan(&cmdline)
		p := &head[0]
		for p != nil {
			if p.cmd == cmdline {
				if p.handler == nil {
					fmt.Println(p.desc)
				} else {
					p.handler()
				}
				break
			}
			p = p.next
		}
		if p == nil {
			fmt.Println("HAHA, VERY FUNNY, DO THAT AGAIN")
		}
	}
}

func Help() {
	fmt.Println("How can I help you?")
	p := &head[0]
	for p != nil && p.next != nil {
		fmt.Println(p.cmd + " - " + p.desc)
		p = p.next
	}
}

func Quit() {
	fmt.Println("OK OK I'm done.")
	os.Exit(0)
}

func Nut() {
	fmt.Println("I could be bounded in a nutshell and count myself a king of infinite space.")
	fmt.Println("https://github.com/phantomT/T-Shell")
}
