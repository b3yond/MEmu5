package main

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/utils"
)

func refreshPersonaWin() {

	windowList[1].(*congo.TWindow).WClear()
	//player = *objectList[0].(*TPersona)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Alias: "+player.GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Device type: "+player.device.deviceType, congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Device model: "+player.device.model, congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona User Mode: "+player.simSence, congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Grid: "+player.grid.GetGridName(), congo.ColorGreen)
	if checkLinkLock(player) == true {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("WARNING: LINK-LOCK DETECTED!", congo.ColorRed)
	}
	if player.GetPhysicalLocation() == true {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("WARNING: Physical location tracked!", congo.ColorYellow)
	}
	//device := player.GetDevice()
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Running Programs: ", congo.ColorGreen)
	var rPrgLst []string
	loadedPrgsQty := 0
	for j := range player.GetDeviceSoft().programStatus {
		//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Program: "+player.GetDeviceSoft().programName[j]+" | Status: "+player.GetDeviceSoft().programStatus[j], congo.ColorGreen)
		if player.GetDeviceSoft().programStatus[j] == "Running" {
			rPrgLst = append(rPrgLst, player.GetDeviceSoft().programName[j])
			loadedPrgsQty++
		}
	}
	if loadedPrgsQty > player.GetMaxRunningPrograms() {
		player.CrashRandomProgram()
	}
	for i := 0; i < player.GetMaxRunningPrograms(); i++ {
		if len(rPrgLst) < player.GetMaxRunningPrograms() {
			rPrgLst = append(rPrgLst, "--EMPTY--")
		}
		congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Slot "+strconv.Itoa(i+1)+": "+rPrgLst[i], congo.ColorGreen)
	}
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Total loaded programs: "+strconv.Itoa(loadedPrgsQty), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Attribute Array: ", congo.ColorGreen)
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn(strconv.Itoa(player.GetAttack())+" "+strconv.Itoa(player.GetSleaze())+" "+strconv.Itoa(player.GetDataProcessing())+" "+strconv.Itoa(player.GetFirewall()), congo.ColorGreen)
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona MCM: " + strconv.Itoa(objectList[0].(IPersona).GetMatrixCM()), congo.ColorYellow)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Att  : "+strconv.Itoa(player.GetAttack()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Slz  : "+strconv.Itoa(player.GetSleaze()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" DtPr : "+strconv.Itoa(player.GetDataProcessing()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Fwll : "+strconv.Itoa(player.GetFirewall()), congo.ColorGreen)
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)
	col := congo.ColorDefault
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Condition Monitor:", congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Matrix  : ", congo.ColorDefault)
	if player.GetMatrixCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("DEVICE BRICKED!!!", congo.ColorRed)

	} else {
		for i := 0; i < player.GetMatrixCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorGreen)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}

		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Stun    : ", congo.ColorYellow)
	if player.GetStunCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("UNCONSCIOUS!!!", congo.ColorRed)
		//congo.WindowsMap.ByTitle["Log"].WPrint("Emergency logging  terminated...", congo.ColorGreen)
	} else {
		for i := 0; i < player.GetStunCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorYellow)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}
		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Physical: ", congo.ColorYellow)

	if player.GetPhysCM() < player.GetBody() {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("DEAD!!!", congo.ColorRed)
	} else if player.GetPhysCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("CRITICAL!!!", congo.ColorRed)
	} else {
		for i := 0; i < player.GetPhysCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorRed)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}
		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)

	if player.GetInitiative() > 9000 {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona Initiative: null", congo.ColorRed)
	} else {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona Initiative: "+strconv.Itoa(player.GetInitiative()), congo.ColorYellow)
	}
	if player.IsConnected() == false {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona disconnected...", congo.ColorYellow)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Session terminated...", congo.ColorGreen)

	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Total Objects: "+strconv.Itoa(len(objectList)), congo.ColorYellow)

	totalMarks := player.CountMarks()
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Confirmed Marks on Persona: "+strconv.Itoa(totalMarks), congo.ColorYellow)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Matrix Search in: "+strconv.Itoa(player.GetSearchResultIn()), congo.ColorYellow)
	for i := range player.GetSearchProcess().SearchIconName {
		name := player.searchProcessStatus.SearchIconName[i]
		objType := player.searchProcessStatus.SearchIconType[i]
		timeTotal := player.searchProcessStatus.SearchTime[i]
		timeSpent := player.searchProcessStatus.SpentTurns[i]
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Search: "+objType+" "+name, congo.ColorGreen)
		currentPer := 0
		turnsPart := 0
		if timeSpent != 0 {
			turnsPart = (100 / timeTotal)
			r := player.GetInitiative()/10 + 1
			currentPer = utils.Min((100/timeTotal*(timeSpent))+turnsPart/r-1, 100)
		}
		congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Progress: "+strconv.Itoa(currentPer)+"%", congo.ColorGreen)
	}
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)

	//fow := player.GetFieldOfView()
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn(fmt.Sprintf("FoW: %v", fow), congo.ColorYellow)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("0", congo.ColorDefault)

}

func refreshGridWin() {
	if player.GetMatrixCM() < 1 {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
	}
	windowList[2].(*congo.TWindow).WClear()
	congo.WindowsMap.ByTitle["Grid"].WPrintLn("Grid:", congo.ColorGreen)
	congo.WindowsMap.ByTitle["Grid"].WPrintLn(player.grid.GetGridName(), congo.ColorGreen)
	if player.CheckRunningProgram("Baby Monitor") {
		warningColor := congo.ColorGreen
		if player.grid.GetOverwatchScore() < 20 {
			warningColor = congo.ColorGreen
		} else if player.grid.GetOverwatchScore() < 31 {
			warningColor = congo.ColorYellow
		} else {
			warningColor = congo.ColorRed
		}
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Overwatch Score: "+strconv.Itoa(player.grid.GetOverwatchScore()), warningColor)
	} else {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Overwatch Score: "+strconv.Itoa(player.grid.GetLastSureOS())+" or more", congo.ColorYellow)
	}

	congo.WindowsMap.ByTitle["Grid"].WPrintLn("Host:", congo.ColorGreen)
	host := player.GetHost().name
	if host == "Matrix" {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("--not in Host--", congo.ColorGreen)
	} else {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn(" "+host, congo.ColorYellow)
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Host Alert Status: ", congo.ColorGreen)
		statusColor := congo.ColorDefault
		alert := player.GetHost().GetHostAlertStatus()
		switch alert {
		case "No Alert":
			statusColor = congo.ColorGreen
		case "Passive Alert":
			statusColor = congo.ColorYellow
		case "Active Alert":
			statusColor = congo.ColorRed
		default:
		}
		congo.WindowsMap.ByTitle["Grid"].WPrintLn(" "+alert, statusColor)
	}

}

func refreshEnviromentWin() {
	congo.WindowsMap.ByTitle["Enviroment"].WClear()

	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("ObjByName: %v", ObjByNames), congo.ColorYellow)

	var row string
	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Turn №: "+strconv.Itoa(Turn), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(STime, congo.ColorDefault)
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(generateShadowrunTime(), congo.ColorDefault)

	for o := range gridList {
		if host, ok := gridList[o].(*THost); ok {
			//host := *gridList[o].(*THost)
			var sampleCode [30]string
			sampleCode[0] = "Spotted" //[0]
			sampleCode[1] = "Unknown" //[1]
			var checkFoW [30]string
			checkFoW = sampleCode
			whatCanSee := player.canSee.KnownData[host.GetID()]
			if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Icon: "+host.GetName(), congo.ColorGreen)
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Type: Host", congo.ColorGreen)
			}
			if checkFoW[5] == whatCanSee[5] && checkFoW[5] != "Unknown" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: "+whatCanSee[5], congo.ColorGreen)
			}
			if whatCanSee[5] != "Unknown" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: "+strconv.Itoa(host.GetDeviceRating()), congo.ColorRed)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: Unknown", congo.ColorYellow)
			}
			Att := "Unknown"
			Slz := "Unknown"
			DtPrc := "Unknown"
			Frw := "Unknown"
			if whatCanSee[7] != "Unknown" {
				Att = strconv.Itoa(host.GetAttack())
			}
			if whatCanSee[8] != "Unknown" {
				Slz = strconv.Itoa(host.GetSleaze())
			}
			if whatCanSee[9] != "Unknown" {
				DtPrc = strconv.Itoa(host.GetDataProcessing())
			}
			if whatCanSee[10] != "Unknown" {
				Frw = strconv.Itoa(host.GetFirewall())
			}
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--Host Attribute Array--", congo.ColorGreen)
			//Show Host Attack
			if whatCanSee[7] != "Unknown" {
				Att = strconv.Itoa(host.GetAttack())
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Attack: "+Att, congo.ColorGreen)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Attack: "+Att, congo.ColorYellow)
			}
			//Show Host Sleaze
			if whatCanSee[8] != "Unknown" {
				Att = strconv.Itoa(host.GetSleaze())
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sleaze: "+Slz, congo.ColorGreen)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sleaze: "+Slz, congo.ColorYellow)
			}
			//Show Host DataProcessing
			if whatCanSee[9] != "Unknown" {
				Att = strconv.Itoa(host.GetDataProcessing())
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Data Processing: "+DtPrc, congo.ColorGreen)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Data Processing: "+DtPrc, congo.ColorYellow)
			}
			//Show Host Firewall
			if whatCanSee[10] != "Unknown" {
				Att = strconv.Itoa(host.GetFirewall())
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Firewall: "+Frw, congo.ColorGreen)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Firewall: "+Frw, congo.ColorYellow)
			}
			//Show Host Grid
			if whatCanSee[13] != "Unknown" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Grid: "+host.GetGridName(), congo.ColorGreen)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Grid: Unknown", congo.ColorYellow)
			}
			if host.GetHostAlertStatus() == "No Alert" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorGreen)
			} else if host.GetHostAlertStatus() == "Passive Alert" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorYellow)
			} else if host.GetHostAlertStatus() == "Active Alert" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorRed)
			}
			if whatCanSee[4] != "Unknown" {
				for i := 0; i < host.GetDeviceRating(); i++ {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint(host.icState.icName[i]+": ", congo.ColorGreen)
					if host.icState.icStatus[i] == true {
						congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Active", congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Passive", congo.ColorGreen)
					}
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				}
			}
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("-DEBUG---------------------", congo.ColorGreen)
			for i := 0; i < host.GetDeviceRating(); i++ {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint(host.icState.icName[i]+": ", congo.ColorGreen)
				if host.icState.icStatus[i] == true {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Active", congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Passive", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
			}
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("----------------------", congo.ColorGreen)

			marks := host.GetMarkSet()
			marksOnPlayer := player.GetMarkSet()
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks on Host: %v", marks), congo.ColorYellow)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks on Player: %v", marksOnPlayer), congo.ColorYellow)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------------", congo.ColorGreen)
			//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(host.(*THost).GetName(), congo.ColorRed)
		}

	}
	///////////////////////////////////
	for o := range objectList {
		if objectList[o].(IObj).GetType() == "IC" {
			ic := *objectList[o].(*TIC)
			var sampleCode [30]string
			sampleCode[0] = "Spotted" //[0]
			sampleCode[1] = "Unknown" //[1]
			var checkFoW [30]string
			checkFoW = sampleCode
			whatCanSee := player.canSee.KnownData[ic.GetID()]
			whatKnowAboutHost := player.canSee.KnownData[ic.GetHost().GetID()]
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--------------------------------", congo.ColorDefault)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(ic.GetName()+"actRed: "+strconv.Itoa(ic.actionReady), congo.ColorDefault)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(ic.GetName()+"Init  : "+strconv.Itoa(ic.initiative), congo.ColorDefault)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(ic.GetName()+"MCM   : "+strconv.Itoa(ic.matrixCM), congo.ColorDefault)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--------------------------------", congo.ColorDefault)
			if ic.GetHost().name == player.GetHost().name && whatCanSee[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--------------------------------", congo.ColorDefault)
				if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Name: "+ic.GetName(), congo.ColorGreen)

				}
				/*if whatCanSee[11] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Icon Name: "+ic.GetName(), congo.ColorGreen)
				}*/
				icMCM := " _ "
				if whatCanSee[2] != "foo" {
					icMCM = strconv.Itoa(ic.GetMatrixCM())
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Matrix Condition Monitor: "+icMCM, congo.ColorGreen)
					marks := ic.GetMarkSet()
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks on IC: %v", marks), congo.ColorYellow)
				}

				if whatCanSee[5] != "Unknown" || whatKnowAboutHost[5] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Rating: "+strconv.Itoa(ic.GetDeviceRating()), congo.ColorGreen)
				}
				Att := "Unknown"
				Slz := "Unknown"
				DtPrc := "Unknown"
				Frw := "Unknown"
				showAttArray := false
				if whatCanSee[7] != "Unknown" || whatKnowAboutHost[7] != "Unknown" {
					Att = strconv.Itoa(ic.GetAttack())
					showAttArray = true
				}
				if whatCanSee[8] != "Unknown" || whatKnowAboutHost[8] != "Unknown" {
					Slz = strconv.Itoa(ic.GetSleaze())
					showAttArray = true
				}
				if whatCanSee[9] != "Unknown" || whatKnowAboutHost[9] != "Unknown" {
					DtPrc = strconv.Itoa(ic.GetDataProcessing())
					showAttArray = true
				}
				if whatCanSee[10] != "Unknown" || whatKnowAboutHost[10] != "Unknown" {
					Frw = strconv.Itoa(ic.GetFirewall())
					showAttArray = true
				}
				if showAttArray == true {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("---IC Attribute Array---", congo.ColorGreen)

					//Show Host Attack
					if whatCanSee[7] != "Unknown" || whatKnowAboutHost[7] != "Unknown" {
						Att = strconv.Itoa(ic.GetAttack())
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Attack: "+Att, congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Attack: "+Att, congo.ColorYellow)
					}
					//Show Host Sleaze
					if whatCanSee[8] != "Unknown" || whatKnowAboutHost[8] != "Unknown" {
						Att = strconv.Itoa(ic.GetSleaze())
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Sleaze: "+Slz, congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Sleaze: "+Slz, congo.ColorYellow)
					}
					//Show Host DataProcessing
					if whatCanSee[9] != "Unknown" || whatKnowAboutHost[9] != "Unknown" {
						Att = strconv.Itoa(ic.GetDataProcessing())
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Data Processing: "+DtPrc, congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Data Processing: "+DtPrc, congo.ColorYellow)
					}
					//Show Host Firewall
					if whatCanSee[10] != "Unknown" || whatKnowAboutHost[10] != "Unknown" {
						Att = strconv.Itoa(ic.GetFirewall())
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Firewall: "+Frw, congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Firewall: "+Frw, congo.ColorYellow)
					}
					/*if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Name: "+ic.GetName(), congo.ColorGreen)
					}*/
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------", congo.ColorGreen)
				}
			}
			//	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Name: "+ic.GetName()+"; id "+strconv.Itoa(ic.GetID()), congo.ColorGreen)
			//	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--------------------------------", congo.ColorDefault)
		}
	}

	for o := range objectList {
		pText := make([]string, 0)
		//row = "Object: " + /*getPersonaAlias()*/ strconv.Itoa(o)
		//pText = append(pText, row)

		if obj, ok := objectList[o].(IDevice); ok {
			row = obj.(IIcon).GetName()
			pText = append(pText, row)
			oName := obj.GetModel()
			row = oName + strconv.Itoa(obj.GetID())
			pText = append(pText, row)
			row = "Device Rating: " + strconv.Itoa(obj.GetDeviceRating())
			pText = append(pText, row)
			row = "Matrix Condition Monitor: " + strconv.Itoa(obj.(IDevice).GetMatrixCM())
			pText = append(pText, row)
			row = "------------------------"
			pText = append(pText, row)
			//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("%v - ", MActions.MActionMap), congo.ColorGreen)
			marks := obj.(*TDevice).GetMarkSet()
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks: %v", marks), congo.ColorYellow)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(obj.(IDevice).GetName(), congo.ColorRed)
		}
		for i := range pText {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(pText[i], congo.ColorGreen)

		}
		//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(obj.(IDevice).GetName() , congo.ColorGreen)
	}

	for o := range objectList {
		if obj, ok := objectList[o].(IFile); ok {
			file := obj
			var sampleCode [30]string
			sampleCode[0] = "Spotted" //[0]
			sampleCode[1] = "Unknown" //[1]
			var checkFoW [30]string
			checkFoW = sampleCode
			whatCanSee := player.canSee.KnownData[file.GetID()]
			//testMarks := file.GetMarkSet()
			//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks on "+file.GetName()+" : %v", testMarks), congo.ColorYellow)
			if file.GetHost() == player.GetHost() {
				//if file.GetHost().name == player.GetHost().name {

				if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--------------------------------", congo.ColorDefault)
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Icon: "+file.GetType()+" "+strconv.Itoa(file.GetID()), congo.ColorGreen)
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Name: "+file.GetFileName(), congo.ColorGreen)

					if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						b := file.GetDataBombRating()
						if b > 0 {
							congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: "+strconv.Itoa(file.GetDataBombRating()), congo.ColorYellow)
						} else {
							congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: "+strconv.Itoa(file.GetDataBombRating()), congo.ColorGreen)
						}
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: Unknown", congo.ColorYellow)
					}
					if whatCanSee[12] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						e := file.GetEncryptionRating()
						if e > 0 {
							congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: "+strconv.Itoa(file.GetEncryptionRating()), congo.ColorYellow)
						} else {
							congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: "+strconv.Itoa(file.GetEncryptionRating()), congo.ColorGreen)
						}
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: Unknown", congo.ColorYellow)
					}
					if whatCanSee[15] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Size: "+strconv.Itoa(file.GetSize())+" Mp", congo.ColorGreen)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Size: Unknown", congo.ColorGreen)
					}
					if whatCanSee[1] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Last Edit Time: "+file.GetLastEditDate(), congo.ColorGreen)
					}

					if file.GetSilentRunningMode() == true {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(file.GetName()+" is silent running", congo.ColorRed)
					}
					marks := file.GetMarkSet()
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks: %v", marks), congo.ColorYellow)
					//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------------", congo.ColorDefault)
				}
			}
		}
	}
}

func refreshProcessWin() {
	congo.WindowsMap.ByTitle["Process"].WDraw()

}

func printLog(s string, color congo.TColor) {
	//if SourceIcon != nil {
	//	if SourceIcon.(IObj).GetFaction() == player.GetFaction() { //вылетает при выборе хода - что-то связанное с тем что оно берет объект но получает нил
	congo.WindowsMap.ByTitle["Log"].WPrintLn(s, color)
	hold()
	//	}
	//}
}
