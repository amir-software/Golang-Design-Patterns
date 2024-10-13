package main



type Client struct {
}

func (c *Client) UseTvWith220Voltage(tv Tv) {
    tv.getTvType()
    tv.plugInTvWith220Volt()
}


type Tv interface{
	plugInTvWith220Volt()
    getTvType()
}