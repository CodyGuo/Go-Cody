package main

var (
	loginXML = `<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
  <s:Header>
    <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>{{.Username}}</wsse:Username>
        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">{{.Password}}</wsse:Password>
        <wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">{{.Nonce64}}</wsse:Nonce>
        <wsu:Created>{{.Created}}</wsu:Created>
      </wsse:UsernameToken>
    </wsse:Security>
  </s:Header>
  `
	deviceInfoXML = `<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <GetDeviceInformation xmlns="http://www.onvif.org/ver10/device/wsdl"></GetDeviceInformation>
  </s:Body>
</s:Envelope>`

	networkInterfacesXML = `<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <GetNetworkInterfaces xmlns="http://www.onvif.org/ver10/device/wsdl"></GetNetworkInterfaces>
  </s:Body>
</s:Envelope>`
)

type Envelope struct {
	Body Body `xml:"Body"`
}
type Body struct {
	GetDeviceInformationResponse GetDeviceInformationResponse `xml:"GetDeviceInformationResponse"`
	GetNetworkInterfacesResponse GetNetworkInterfacesResponse
}

type GetDeviceInformationResponse struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces NetworkInterfaces
}

type NetworkInterfaces struct {
	Info Info
}

type Info struct {
	Name      string
	HwAddress string
}
