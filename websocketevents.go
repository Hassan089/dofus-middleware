package main

import (
	"dofusmiddleware/database"
	"dofusmiddleware/options"
	"dofusmiddleware/socket"
	"dofusmiddleware/windowmanagement"
	"dofusmiddleware/world"
	"fmt"
	"strconv"
	"strings"
)

var MapDebug world.Map

func OnSocketConnexion() {
	fmt.Println("[websocket] new connexion.")
	fmt.Println("Send characters...")
	web.SendCharacters(world.Players)
	fmt.Println("Send options...")
	web.SendOptions(options.Options)

	//web.SendMap(MapDebug) // TODO remove
	//web.SendPath(world.Visited) // TODO remove

	randomCharacter := world.GetAConnectedPlayer()
	if randomCharacter != nil {
		fmt.Println("send map...")
		themap := database.GetMap(world.GetAConnectedPlayer().MapId)
		web.SendMap(themap)
	}
}

func OnFocus(args []string) {
	go windowmanagement.SwitchToCharacter(args[1])
}

func OnSetOption(args []string) {
	fmt.Println("websocket input : SET_OPTIONS " + args[1] + " " + args[2])
	if args[1] == "ShowInputPackets" {
		options.Options.ShowInputPackets = args[2] == "true"
	}
	if args[1] == "ShowOutputPackets" {
		options.Options.ShowOutputPackets = args[2] == "true"
	}
	if args[1] == "DispatchMoves" {
		options.Options.DispatchMoves = args[2] == "true"
	}
	if args[1] == "FocusWindowOnCharacterTurn" {
		options.Options.FocusWindowOnCharacterTurn = args[2] == "true"
	}
	if args[1] == "AutoJoinFight" {
		options.Options.AutoJoinFight = args[2] == "true"
	}
	if args[1] == "AutoReadyFight" {
		options.Options.AutoReadyFight = args[2] == "true"
	}
	web.SendOptions(options.Options)
}

func OnSetCharacterOption(args []string) {
	fmt.Println("websocket input : SET_CHARACTER_OPTIONS " + args[1] + " " + args[2] + " " + args[3])
	if args[2] == "OptionAutoPassTurn" {
		world.GetPlayer(args[1]).OptionAutoPassTurn = args[3] == "true"
	}
	if args[2] == "OptionAutoFight" {
		world.GetPlayer(args[1]).OptionAutoFight = args[3] == "true"
	}
	if args[2] == "OptionAutoStartFight" {
		world.GetPlayer(args[1]).OptionAutoStartFight = args[3] == "true"
	}
	web.SendCharacters(world.Players)
}

func OnMoveToMapInstruction(args []string) {
	idPlayer := args[1]
	idMap, _ := strconv.Atoi(args[2])
	player := world.GetPlayer(idPlayer)
	if idMap == 0 {
		fmt.Println("[websocket] cant find map id", idMap)
		return
	}
	AddMoveTo(player.IdCharDofus, player.MapId, idMap)
	processMoveTo(*player)
}

func OnMoveToMapPositionInstruction(args []string) {
	idPlayer := args[1]
	position := args[2]
	idMap := database.GetMapIdFromPosition(position)
	if idMap != 0 {
		player := world.GetPlayer(idPlayer)
		AddMoveTo(player.IdCharDofus, player.MapId, idMap)
		processMoveTo(*player)
	}
}

func OnProcessPath(args []string) {

	idMap, _ := strconv.Atoi(args[1])
	cellStart, _ := strconv.Atoi(args[2])
	cellEnd, _ := strconv.Atoi(args[3])

	fmt.Println("process path", idMap, cellStart, cellEnd)

	themap := database.GetMap(idMap)
	path := world.AStar(themap, nil, cellStart, cellEnd)
	encodedPath := world.EncodePath(themap, path)

	web.SendPath(path)
	socket.SendMovePacket(*world.GetAConnectedPlayer().Connexion, encodedPath)
}

func OnSocketMessage(packet string) {
	fmt.Println("[websocket] client sent message", packet)

	parts := strings.Split(packet, "|")
	typepacket := parts[0]
	if typepacket == "FOCUS" {
		OnFocus(parts)
	}
	if typepacket == "SET_OPTION" {
		OnSetOption(parts)
	}
	if typepacket == "SET_CHARACTER_OPTION" {
		OnSetCharacterOption(parts)
	}
	if typepacket == "PROCESS_PATH" {
		OnProcessPath(parts)
	}
	if typepacket == "MOVE_PLAYER_TO_MAP_ID" {
		OnMoveToMapInstruction(parts)
	}
	if typepacket == "MOVE_PLAYER_TO_MAP_POSITION" {
		OnMoveToMapPositionInstruction(parts)
	}
}
