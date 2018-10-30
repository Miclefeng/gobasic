package main

import "fmt"

func main()  {
	rooms := []map[string]string{{"Id": "room1", "name": "name1", "addr": "addr1"}, {"Id": "room2", "name": "name2", "addr": "addr2"}}
	record := []map[string]string{{"roomId": "room1"}, {"roomId": "room2"}}

	for _, v := range record {
		for j, room := range rooms {
			fmt.Println(j, room["Id"])
			if v["roomId"] == room["Id"] {
				v["name"] = room["name"]
				v["addr"] = room["addr"]
				break
			}
		}
	}
	fmt.Println()
	existRoom := make(map[string]map[string]string, 2)
	for _, v := range record {
		if _, ok := existRoom[v["roomId"]]["name"]; ok {
			v["name"] = existRoom[v["roomId"]]["name"]
			v["addr"] = existRoom[v["roomId"]]["addr"]
			break
		}
		for m, room := range rooms {
			fmt.Println(m, room["Id"])
			if v["roomId"] == room["Id"] {
				v["name"] = room["name"]
				v["addr"] = room["addr"]
				// 初始化嵌套的map
				if existRoom[v["roomId"]] == nil {
					existRoom[v["roomId"]] = make(map[string]string)
				}
				existRoom[v["roomId"]]["name"] = room["name"]
				existRoom[v["roomId"]]["addr"] = room["addr"]
				break
			}
		}
	}
}
