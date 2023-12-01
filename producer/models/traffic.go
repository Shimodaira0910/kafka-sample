package models

type Traffic struct{
	Type string `json:"type"`
	ID int `json:"id"`
	Nodes []int `json:"nodes"`
	Tags string `json:"tags"`
}