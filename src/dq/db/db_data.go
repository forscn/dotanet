package db

type DB_CharacterInfo struct {
	Playerid   int32  `json:"playerid"`
	Uid        int32  `json:"uid"`
	Name       string `json:"name"`
	Typeid     int32  `json:"typeid"`
	Level      int32  `json:"level"`
	Experience int32  `json:"experience"`
	Gold       int32  `json:"gold"`
}