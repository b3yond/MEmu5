package main

import (
	//"strconv"

	"strings"

	"github.com/Galdoba/ConGo/congo"
	//"encoding/base64"
	"encoding/hex"
)

//UserInput -
func UserInput(input string) bool {
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("Processing: '"+input+"'", congo.ColorGreen)
	var mActionName string
	var actionIsGood bool
	var comm []string

	command = input
	command = formatString(command)
	comm = strings.SplitN(command, ">", 6)
	text := formatString(input)
	text = cleanText(text)
	congo.WindowsMap.ByTitle["Log"].WPrintLn(text, congo.ColorGreen)
	hold()

	if len(comm) < 2 {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! Sintax Error!", congo.ColorRed)
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET]' Format", congo.ColorDefault)
		congo.WindowsMap.ByTitle["Log"].WDraw()
		return false
	}
	//////Checking if action isValid
	mAction := comm[1]
	mAction = formatString(mAction)
	mAction = cleanText(mAction)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(mAction, congo.ColorGreen)
	//printLog(mAction, congo.ColorRed)
	mActionName = mAction
	//printLog(mActionName, congo.ColorRed)
	actionIsGood, mActionName = checkAction(mAction)
	if actionIsGood == true {
	//	congo.WindowsMap.ByTitle["Log"].WPrintLn("Action: "+mActionName+" is correct", congo.ColorYellow)
	}
	checkSource(comm[0])
	if mActionName == "EXIT_HOST" || mActionName == "ERASE_MARK" || mActionName == "CHECK_OVERWATCH_SCORE" || mActionName == "LONGACT" || mActionName == "SWITCH_INTERFACE_MODE" {
		TargetIcon = SourceIcon
		doAction(mActionName)
		return true
	}
	if mActionName == "WAIT" || mActionName == "FULL_DEFENCE" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		doAction(mActionName)
		return true
	}

	if mActionName == "SWAP_ATTRIBUTES" || mActionName == "SWAP_PROGRAMS" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		if len(comm) < 4 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [ATTRIBUTE 1] > [ATTRIBUTE 2]' Format", congo.ColorDefault)
			return false
		} else if len(comm) == 4 {
			//TargetIcon = text
			TargetIcon = SourceIcon
			doAction(mActionName)
			return true
		}
		return false
	}

	if mActionName == "LOAD_PROGRAM" || mActionName == "UNLOAD_PROGRAM" || mActionName == "LOGIN" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		if len(comm) < 3 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [ATTRIBUTE 1] > [ATTRIBUTE 2]' Format", congo.ColorDefault)
			return false
		} else if len(comm) == 3 {
			//TargetIcon = text
			TargetIcon = SourceIcon
			doAction(mActionName)
			return true
		}
		return false
	}

	if mActionName == "MATRIX_PERCEPTION" {
		if len(comm) > 2 {
			target := comm[2]
			target = formatString(target)
			target = cleanText(target)
			if target == "ALL" {
				mActionName = "SCAN_ENVIROMENT"
				//TargetIcon = "ALL"
				doAction(mActionName)
				return true
			}
		}
	}

	if mActionName == "MATRIX_SEARCH" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		if len(comm) < 3 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET TYPE] > [TARGET NAME]' Format", congo.ColorDefault)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("[TARGET NAME] is optional, if left blank random name will be generated", congo.ColorDefault)
			return false
		}
		doAction(mActionName)
		return true
	}
	if len(comm) < 3 {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! Sintax Error!", congo.ColorRed)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET]' Format", congo.ColorDefault)
		congo.WindowsMap.ByTitle["Log"].WDraw()
		return false
	}

	if checkTarget(comm[2], mActionName) == true {
		doAction(mActionName)
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Unknown target", congo.ColorGreen)

	}

	congo.WindowsMap.ByTitle["Log"].WDraw()

	return true
}

func formatString(s string) string {
	s = strings.ToUpper(s)
	s = strings.Replace(s, " ", "_", -1)
	//s = strings.Replace(s, "-2M", "-2m", -1)
	//s = strings.Replace(s, "-3M", "-3m", -1)
	return s
}

func checkSource(source string) bool {
	source = formatString(source)
	source = cleanText(source)
	isGood := false
	//var alias string
	for _, obj := range ObjByNames {
		if srcObj, ok := obj.(IPersona); ok {
			//if srcObj.(IIcon).GetType() == "Persona" {
			alias := string(srcObj.(IPersona).GetName())
			alias = formatString(alias)
			s := (hex.EncodeToString([]byte(source)))
			a := (hex.EncodeToString([]byte(alias)))
			if a == s {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("SourceIcon is " + objectList[i].(*TPersona).GetName(), congo.ColorYellow)
				if obj.(IPersona).IsPlayer() == true {
					SourceIcon = srcObj
					isGood = true
					return isGood
				}

			}
		}

	}

	for i := range objectList {
		//srcObj := objectList[i]
		if srcObj, ok := objectList[i].(IPersona); ok {
			//if srcObj.(IIcon).GetType() == "Persona" {
			alias := string(srcObj.(IPersona).GetName())
			alias = formatString(alias)
			s := (hex.EncodeToString([]byte(source)))
			a := (hex.EncodeToString([]byte(alias)))
			if a == s {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("SourceIcon is " + objectList[i].(*TPersona).GetName(), congo.ColorYellow)
				if objectList[i].(IPersona).IsPlayer() == true {
					SourceIcon = objectList[i]
					isGood = true
					return isGood
				}

			}
		}
	}
	return isGood
}

func checkTarget(target, mActionName string) bool {
	if pickTarget(target, mActionName) {
		return true
	}
	//printLog("Error: target not found", congo.ColorGreen)
	return false
}

func pickHost(target, mActionName string) bool { //ненужная функция?
	for i := range hostList {
		trgObj := objectList[i]

		if trgObj.(IObj).GetType() == "Host" {
			alias := trgObj.(IObj).GetName()
			alias = formatString(alias)
			alias = cleanText(alias)
			if alias == target {

			}
		}
	}
	return false
}

func pickTarget(target, mActionName string) bool {
	target = formatString(target)
	target = cleanText(target)
	for _, obj := range ObjByNames {
		if grid, ok := obj.(IGrid); ok {
			var alias string
			alias = grid.GetName()
			alias = formatString(alias)
			alias = cleanText(alias)
			if alias == target {
				TargetIcon = grid
				return true
			}
		}
		if icon, ok := obj.(IIcon); ok {
			var alias string
			alias = icon.GetName()
			alias = formatString(alias)
			alias = cleanText(alias)
			if alias == target {
				TargetIcon = icon
				return true
			}
		}
	}
	return false
}

/*func checkValidTarget(mActionName string) (trgtType []string, valid bool) { //тоже не нужная?
	//var trgtType []string

	switch mActionName {
	case "BRUTE_FORCE":
		trgtType = append(trgtType, "Icon")
		trgtType = append(trgtType, "Host")
		trgtType = append(trgtType, "Grid")
		trgtType = append(trgtType, "File")
		trgtType = append(trgtType, "IC")
		valid = true
	case "CRACK_FILE":
		trgtType = append(trgtType, "File")
		valid = true
	case "CHECK_OVERWATCH_SCORE":
		trgtType = append(trgtType, "Grid")
		valid = true
	case "DATA_SPIKE":
		trgtType = append(trgtType, "Persona")
		trgtType = append(trgtType, "Icon")
		trgtType = append(trgtType, "IC")
		valid = true
	case "DISARM_DATABOMB":
		trgtType = append(trgtType, "File")
		valid = true
	case "EDIT":
		trgtType = append(trgtType, "File")
		valid = true
	case "ENTER_HOST":
		trgtType = append(trgtType, "Host")
		valid = true
	case "EXIT_HOST":
		trgtType = append(trgtType, "Host")
		valid = true
	case "ERASE_MARK":
		valid = true
	case "GRID_HOP":
		trgtType = append(trgtType, "Grid")
		valid = true
	case "HACK_ON_THE_FLY":
		trgtType = append(trgtType, "Icon")
		trgtType = append(trgtType, "Host")
		trgtType = append(trgtType, "Grid")
		trgtType = append(trgtType, "File")
		trgtType = append(trgtType, "IC")
		valid = true
	case "MATRIX_PERCEPTION":
		trgtType = append(trgtType, "Icon")
		trgtType = append(trgtType, "Host")
		trgtType = append(trgtType, "Grid")
		trgtType = append(trgtType, "File")
		trgtType = append(trgtType, "IC")
		valid = true
	case "SCAN_ENVIROMENT":

		valid = true

	case "MATRIX_SEARCH":
		//всегда валидно ибо цели нет

		valid = true
	case "SWAP_ATTRIBUTE":
		valid = true
	case "LOAD_PROGRAM":
		valid = true
	case "UNLOAD_PROGRAM":
		valid = true
	case "SET_DATABOMB":
		trgtType = append(trgtType, "File")
		valid = true
	case "SWAP_PROGRAMS":
		valid = true
	case "SWITCH_INTERFACE_MODE":
		valid = true
	case "LONGACT":
		//trgtType = append(trgtType, "Host")
		valid = true
	default:
		trgtType = append(trgtType, "NO_VALID")
		//trgtType[0] = "noValidTarget"
		valid = false
		//return trgtType
	}
	return trgtType, valid
}*/

func cleanText(s string) string {
	out := ""
	plain := hex.EncodeToString([]byte(s))
	//plain = strings.Replace(plain, "10", "", -1)
	char := strings.Split(plain, "1001")
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(plain, congo.ColorGreen)
	for i := range char {

		if char[i] == "10" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("remove \x10")
		} else if char[i] == "01" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("remove \x01", congo.ColorYellow)
		} else {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("keep x__")
			out = out + char[i]
		}
	}
	hexOut, _ := hex.DecodeString(out)

	return string(hexOut)

}
