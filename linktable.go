package main

/********************************************************************************************/
/* Copyright (C), SSE@USTC, 2021-2022                                                       */
/*                                                                                          */
/*  FILE NAME             :  linktable.go                                                   */
/*  PRINCIPAL AUTHOR      :  Terabyte                                                       */
/*  SUBSYSTEM NAME        :  LinkTable                                                      */
/*  MODULE NAME           :  LinkTable                                                      */
/*  LANGUAGE              :  Go                                                             */
/*  TARGET ENVIRONMENT    :  Any                                                            */
/*  DATE OF FIRST RELEASE :  2022/04/06                                                     */
/*  DESCRIPTION           :  Interface of Link Table                                        */
/********************************************************************************************/

/*
 * Revision log:
 *
 * Add Basic Interface of Link Table, 2022/04/06
 * Created by Terabyte, 2022/04/06
 * Add Callback interface, 2022/06/22
 *
 */

import (
	"fmt"
	"sync"
)

const SUCCESS = 0
const FAILURE = -1

/*
 * LinkTable Node Type
 */
type LinkTableNode struct {
	pNext *LinkTableNode
}

/*
 * Define LinkTable
 */
type LinkTable struct {
	pHead     *LinkTableNode
	pTail     *LinkTableNode
	SumOfNode int
	mutex     sync.Mutex
}

/*
 * Create a LinkTable
 */
func CreateLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	if pLinkTable == nil {
		return nil
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return pLinkTable
}

/*
 * Delete a LinkTable
 */
func DeleteLinkTable(pLinkTable *LinkTable) int {
	if pLinkTable == nil {
		return FAILURE
	}
	for pLinkTable.pHead != nil {
		var p *LinkTableNode = pLinkTable.pHead
		pLinkTable.mutex.Lock()
		pLinkTable.pHead = p.pNext
		pLinkTable.SumOfNode--
		pLinkTable.mutex.Unlock()
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return SUCCESS
}

/*
 * Add a LinkTableNode to LinkTable
 */
func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return FAILURE
	}
	pNode.pNext = nil
	pLinkTable.mutex.Lock()
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
	}
	if pLinkTable.pTail == nil {
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.SumOfNode++
	pLinkTable.mutex.Unlock()
	return SUCCESS
}

/*
 * Delete a LinkTableNode from LinkTable
 */
func DelLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return FAILURE
	}
	pLinkTable.mutex.Lock()
	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.pNext
		pLinkTable.SumOfNode--
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		pLinkTable.mutex.Unlock()
		return SUCCESS
	}
	var pTempNode *LinkTableNode = pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode.pNext == pNode {
			pTempNode.pNext = pTempNode.pNext.pNext
			pLinkTable.SumOfNode--
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			pLinkTable.mutex.Unlock()
			return SUCCESS
		}
		pTempNode = pTempNode.pNext
	}
	pLinkTable.mutex.Unlock()
	return FAILURE
}

/*
 * Search a LinkTableNode from LinkTable
 * Condition func(pNode *LinkTableNode, args interface{}) int
 */
func SearchLinkTableNode(pLinkTable *LinkTable, Condition func(pNode *LinkTableNode, args interface{}) int, args interface{}) *LinkTableNode {
	if pLinkTable == nil || Condition == nil {
		return nil
	}
	var pNode *LinkTableNode = pLinkTable.pHead
	for pNode != nil {
		if Condition(pNode, args) == SUCCESS {
			return pNode
		}
		pNode = pNode.pNext
	}
	return nil
}

/*
 * get LinkTableHead
 */
func GetLinkTableHead(pLinkTable *LinkTable) *LinkTableNode {
	if pLinkTable == nil {
		fmt.Println("LinkTable is empty")
		return nil
	}
	return pLinkTable.pHead
}

/*
 * get next LinkTableNode
 */
func GetNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode {
	if pLinkTable == nil || pNode == nil {
		fmt.Println("Linktable is empty")
		return nil
	}
	var pTempNode *LinkTableNode = pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.pNext
		}
		pTempNode = pTempNode.pNext
	}
	return nil
}
