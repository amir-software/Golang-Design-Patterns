package main

func main (){
	client1 := Client{}

	tv220 := Tv220{name: "Tv with 220 Voltage required"}
	Tv110 := Tv110{name: "Tv with 110 Voltage required"}

	client1.UseTvWith220Voltage(&tv220)

	adapter := Adapt110to220{tv110 : &Tv110}

	client1.UseTvWith220Voltage(&adapter)
}