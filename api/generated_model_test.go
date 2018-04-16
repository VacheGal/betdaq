package api

import (
	"fmt"
	"testing"

	"github.com/james-wilder/betdaq/soap"
)

func TestEncode(t *testing.T) {
	req := &GetPrices{
		GetPricesRequest: GetPricesRequest{
			ThresholdAmount:             "0",
			NumberForPricesRequired:     -1,
			NumberAgainstPricesRequired: -1,
			MarketIds: []int64{
				483492,
			},
		},
	}
	data, err := soap.Encode(&req, "username", "xxx")
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(data))
	if string(data) != expectedSoapRequest {
		t.Log("Not expected SOAP request output")
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	t.Run("Test GetOddsLadder response", func(t *testing.T) {
		var resp GetOddsLadderResponse
		err := soap.Decode([]byte(actualGetLadderResponse), &resp)
		fmt.Println(resp)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if len(resp.GetOddsLadderResult.ReturnStatus) == 0 {
			t.Fail()
			return
		}
		fmt.Println(resp.GetOddsLadderResult.ReturnStatus[0].CallId)
		if resp.GetOddsLadderResult.ReturnStatus[0].CallId != "26091ffa-e9e7-437a-aaf5-6e690bc3e33a" {
			t.Fail()
		}
		fmt.Println("Ladders:", len(resp.GetOddsLadderResult.Ladder))
		if len(resp.GetOddsLadderResult.Ladder) != 495 {
			t.Fail()
		}
		fmt.Println("Ladders[3]:", resp.GetOddsLadderResult.Ladder[3])
		if resp.GetOddsLadderResult.Ladder[3].Price != "1.04" {
			t.Fail()
		}
		if resp.GetOddsLadderResult.Ladder[3].Representation != "1.04" {
			t.Fail()
		}
	})
}

var expectedSoapRequest = `<?xml version="1.0" encoding="UTF-8"?>
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <ExternalApiHeader xmlns="http://www.GlobalBettingExchange.com/ExternalAPI/" version="2" languageCode="en" username="username" password="xxx"></ExternalApiHeader>
  </Header>
  <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <GetPrices xmlns="http://www.GlobalBettingExchange.com/ExternalAPI/">
      <getPricesRequest ThresholdAmount="0" NumberForPricesRequired="-1" NumberAgainstPricesRequired="-1" WantMarketMatchedAmount="false" WantSelectionsMatchedAmounts="false" WantSelectionMatchedDetails="false">
        <MarketIds>483492</MarketIds>
      </getPricesRequest>
    </GetPrices>
</Body>
</Envelope>`

var actualGetLadderResponse = `<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <soap:Body>
    <GetOddsLadderResponse xmlns="http://www.GlobalBettingExchange.com/ExternalAPI/">
      <GetOddsLadderResult>
        <ReturnStatus Code="0" Description="Success" CallId="26091ffa-e9e7-437a-aaf5-6e690bc3e33a" />
        <Timestamp>2018-04-12T15:31:45.8388331+00:00</Timestamp>
        <Ladder price="1.01" representation="1.01" />
        <Ladder price="1.02" representation="1.02" />
        <Ladder price="1.03" representation="1.03" />
        <Ladder price="1.04" representation="1.04" />
        <Ladder price="1.05" representation="1.05" />
        <Ladder price="1.06" representation="1.06" />
        <Ladder price="1.07" representation="1.07" />
        <Ladder price="1.08" representation="1.08" />
        <Ladder price="1.09" representation="1.09" />
        <Ladder price="1.1" representation="1.1" />
        <Ladder price="1.11" representation="1.11" />
        <Ladder price="1.12" representation="1.12" />
        <Ladder price="1.13" representation="1.13" />
        <Ladder price="1.14" representation="1.14" />
        <Ladder price="1.15" representation="1.15" />
        <Ladder price="1.16" representation="1.16" />
        <Ladder price="1.17" representation="1.17" />
        <Ladder price="1.18" representation="1.18" />
        <Ladder price="1.19" representation="1.19" />
        <Ladder price="1.2" representation="1.2" />
        <Ladder price="1.21" representation="1.21" />
        <Ladder price="1.22" representation="1.22" />
        <Ladder price="1.23" representation="1.23" />
        <Ladder price="1.24" representation="1.24" />
        <Ladder price="1.25" representation="1.25" />
        <Ladder price="1.26" representation="1.26" />
        <Ladder price="1.27" representation="1.27" />
        <Ladder price="1.28" representation="1.28" />
        <Ladder price="1.29" representation="1.29" />
        <Ladder price="1.3" representation="1.3" />
        <Ladder price="1.31" representation="1.31" />
        <Ladder price="1.32" representation="1.32" />
        <Ladder price="1.33" representation="1.33" />
        <Ladder price="1.34" representation="1.34" />
        <Ladder price="1.35" representation="1.35" />
        <Ladder price="1.36" representation="1.36" />
        <Ladder price="1.37" representation="1.37" />
        <Ladder price="1.38" representation="1.38" />
        <Ladder price="1.39" representation="1.39" />
        <Ladder price="1.4" representation="1.4" />
        <Ladder price="1.41" representation="1.41" />
        <Ladder price="1.42" representation="1.42" />
        <Ladder price="1.43" representation="1.43" />
        <Ladder price="1.44" representation="1.44" />
        <Ladder price="1.45" representation="1.45" />
        <Ladder price="1.46" representation="1.46" />
        <Ladder price="1.47" representation="1.47" />
        <Ladder price="1.48" representation="1.48" />
        <Ladder price="1.49" representation="1.49" />
        <Ladder price="1.5" representation="1.5" />
        <Ladder price="1.51" representation="1.51" />
        <Ladder price="1.52" representation="1.52" />
        <Ladder price="1.53" representation="1.53" />
        <Ladder price="1.54" representation="1.54" />
        <Ladder price="1.55" representation="1.55" />
        <Ladder price="1.56" representation="1.56" />
        <Ladder price="1.57" representation="1.57" />
        <Ladder price="1.58" representation="1.58" />
        <Ladder price="1.59" representation="1.59" />
        <Ladder price="1.6" representation="1.6" />
        <Ladder price="1.61" representation="1.61" />
        <Ladder price="1.62" representation="1.62" />
        <Ladder price="1.63" representation="1.63" />
        <Ladder price="1.64" representation="1.64" />
        <Ladder price="1.65" representation="1.65" />
        <Ladder price="1.66" representation="1.66" />
        <Ladder price="1.67" representation="1.67" />
        <Ladder price="1.68" representation="1.68" />
        <Ladder price="1.69" representation="1.69" />
        <Ladder price="1.7" representation="1.7" />
        <Ladder price="1.71" representation="1.71" />
        <Ladder price="1.72" representation="1.72" />
        <Ladder price="1.73" representation="1.73" />
        <Ladder price="1.74" representation="1.74" />
        <Ladder price="1.75" representation="1.75" />
        <Ladder price="1.76" representation="1.76" />
        <Ladder price="1.77" representation="1.77" />
        <Ladder price="1.78" representation="1.78" />
        <Ladder price="1.79" representation="1.79" />
        <Ladder price="1.8" representation="1.8" />
        <Ladder price="1.81" representation="1.81" />
        <Ladder price="1.82" representation="1.82" />
        <Ladder price="1.83" representation="1.83" />
        <Ladder price="1.84" representation="1.84" />
        <Ladder price="1.85" representation="1.85" />
        <Ladder price="1.86" representation="1.86" />
        <Ladder price="1.87" representation="1.87" />
        <Ladder price="1.88" representation="1.88" />
        <Ladder price="1.89" representation="1.89" />
        <Ladder price="1.9" representation="1.9" />
        <Ladder price="1.91" representation="1.91" />
        <Ladder price="1.92" representation="1.92" />
        <Ladder price="1.93" representation="1.93" />
        <Ladder price="1.94" representation="1.94" />
        <Ladder price="1.95" representation="1.95" />
        <Ladder price="1.96" representation="1.96" />
        <Ladder price="1.97" representation="1.97" />
        <Ladder price="1.98" representation="1.98" />
        <Ladder price="1.99" representation="1.99" />
        <Ladder price="2" representation="2" />
        <Ladder price="2.02" representation="2.02" />
        <Ladder price="2.04" representation="2.04" />
        <Ladder price="2.06" representation="2.06" />
        <Ladder price="2.08" representation="2.08" />
        <Ladder price="2.1" representation="2.1" />
        <Ladder price="2.12" representation="2.12" />
        <Ladder price="2.14" representation="2.14" />
        <Ladder price="2.16" representation="2.16" />
        <Ladder price="2.18" representation="2.18" />
        <Ladder price="2.2" representation="2.2" />
        <Ladder price="2.22" representation="2.22" />
        <Ladder price="2.24" representation="2.24" />
        <Ladder price="2.26" representation="2.26" />
        <Ladder price="2.28" representation="2.28" />
        <Ladder price="2.3" representation="2.3" />
        <Ladder price="2.32" representation="2.32" />
        <Ladder price="2.34" representation="2.34" />
        <Ladder price="2.36" representation="2.36" />
        <Ladder price="2.38" representation="2.38" />
        <Ladder price="2.4" representation="2.4" />
        <Ladder price="2.42" representation="2.42" />
        <Ladder price="2.44" representation="2.44" />
        <Ladder price="2.46" representation="2.46" />
        <Ladder price="2.48" representation="2.48" />
        <Ladder price="2.5" representation="2.5" />
        <Ladder price="2.52" representation="2.52" />
        <Ladder price="2.54" representation="2.54" />
        <Ladder price="2.56" representation="2.56" />
        <Ladder price="2.58" representation="2.58" />
        <Ladder price="2.6" representation="2.6" />
        <Ladder price="2.62" representation="2.62" />
        <Ladder price="2.64" representation="2.64" />
        <Ladder price="2.66" representation="2.66" />
        <Ladder price="2.68" representation="2.68" />
        <Ladder price="2.7" representation="2.7" />
        <Ladder price="2.72" representation="2.72" />
        <Ladder price="2.74" representation="2.74" />
        <Ladder price="2.76" representation="2.76" />
        <Ladder price="2.78" representation="2.78" />
        <Ladder price="2.8" representation="2.8" />
        <Ladder price="2.82" representation="2.82" />
        <Ladder price="2.84" representation="2.84" />
        <Ladder price="2.86" representation="2.86" />
        <Ladder price="2.88" representation="2.88" />
        <Ladder price="2.9" representation="2.9" />
        <Ladder price="2.92" representation="2.92" />
        <Ladder price="2.94" representation="2.94" />
        <Ladder price="2.96" representation="2.96" />
        <Ladder price="2.98" representation="2.98" />
        <Ladder price="3" representation="3" />
        <Ladder price="3.05" representation="3.05" />
        <Ladder price="3.1" representation="3.1" />
        <Ladder price="3.15" representation="3.15" />
        <Ladder price="3.2" representation="3.2" />
        <Ladder price="3.25" representation="3.25" />
        <Ladder price="3.3" representation="3.3" />
        <Ladder price="3.35" representation="3.35" />
        <Ladder price="3.4" representation="3.4" />
        <Ladder price="3.45" representation="3.45" />
        <Ladder price="3.5" representation="3.5" />
        <Ladder price="3.55" representation="3.55" />
        <Ladder price="3.6" representation="3.6" />
        <Ladder price="3.65" representation="3.65" />
        <Ladder price="3.7" representation="3.7" />
        <Ladder price="3.75" representation="3.75" />
        <Ladder price="3.8" representation="3.8" />
        <Ladder price="3.85" representation="3.85" />
        <Ladder price="3.9" representation="3.9" />
        <Ladder price="3.95" representation="3.95" />
        <Ladder price="4" representation="4" />
        <Ladder price="4.1" representation="4.1" />
        <Ladder price="4.2" representation="4.2" />
        <Ladder price="4.3" representation="4.3" />
        <Ladder price="4.4" representation="4.4" />
        <Ladder price="4.5" representation="4.5" />
        <Ladder price="4.6" representation="4.6" />
        <Ladder price="4.7" representation="4.7" />
        <Ladder price="4.8" representation="4.8" />
        <Ladder price="4.9" representation="4.9" />
        <Ladder price="5" representation="5" />
        <Ladder price="5.1" representation="5.1" />
        <Ladder price="5.2" representation="5.2" />
        <Ladder price="5.3" representation="5.3" />
        <Ladder price="5.4" representation="5.4" />
        <Ladder price="5.5" representation="5.5" />
        <Ladder price="5.6" representation="5.6" />
        <Ladder price="5.7" representation="5.7" />
        <Ladder price="5.8" representation="5.8" />
        <Ladder price="5.9" representation="5.9" />
        <Ladder price="6" representation="6" />
        <Ladder price="6.2" representation="6.2" />
        <Ladder price="6.4" representation="6.4" />
        <Ladder price="6.6" representation="6.6" />
        <Ladder price="6.8" representation="6.8" />
        <Ladder price="7" representation="7" />
        <Ladder price="7.2" representation="7.2" />
        <Ladder price="7.4" representation="7.4" />
        <Ladder price="7.6" representation="7.6" />
        <Ladder price="7.8" representation="7.8" />
        <Ladder price="8" representation="8" />
        <Ladder price="8.2" representation="8.2" />
        <Ladder price="8.4" representation="8.4" />
        <Ladder price="8.6" representation="8.6" />
        <Ladder price="8.8" representation="8.8" />
        <Ladder price="9" representation="9" />
        <Ladder price="9.2" representation="9.2" />
        <Ladder price="9.4" representation="9.4" />
        <Ladder price="9.6" representation="9.6" />
        <Ladder price="9.8" representation="9.8" />
        <Ladder price="10" representation="10" />
        <Ladder price="10.5" representation="10.5" />
        <Ladder price="11" representation="11" />
        <Ladder price="11.5" representation="11.5" />
        <Ladder price="12" representation="12" />
        <Ladder price="12.5" representation="12.5" />
        <Ladder price="13" representation="13" />
        <Ladder price="13.5" representation="13.5" />
        <Ladder price="14" representation="14" />
        <Ladder price="14.5" representation="14.5" />
        <Ladder price="15" representation="15" />
        <Ladder price="15.5" representation="15.5" />
        <Ladder price="16" representation="16" />
        <Ladder price="16.5" representation="16.5" />
        <Ladder price="17" representation="17" />
        <Ladder price="17.5" representation="17.5" />
        <Ladder price="18" representation="18" />
        <Ladder price="18.5" representation="18.5" />
        <Ladder price="19" representation="19" />
        <Ladder price="19.5" representation="19.5" />
        <Ladder price="20" representation="20" />
        <Ladder price="21" representation="21" />
        <Ladder price="22" representation="22" />
        <Ladder price="23" representation="23" />
        <Ladder price="24" representation="24" />
        <Ladder price="25" representation="25" />
        <Ladder price="26" representation="26" />
        <Ladder price="27" representation="27" />
        <Ladder price="28" representation="28" />
        <Ladder price="29" representation="29" />
        <Ladder price="30" representation="30" />
        <Ladder price="31" representation="31" />
        <Ladder price="32" representation="32" />
        <Ladder price="33" representation="33" />
        <Ladder price="34" representation="34" />
        <Ladder price="35" representation="35" />
        <Ladder price="36" representation="36" />
        <Ladder price="37" representation="37" />
        <Ladder price="38" representation="38" />
        <Ladder price="39" representation="39" />
        <Ladder price="40" representation="40" />
        <Ladder price="41" representation="41" />
        <Ladder price="42" representation="42" />
        <Ladder price="43" representation="43" />
        <Ladder price="44" representation="44" />
        <Ladder price="45" representation="45" />
        <Ladder price="46" representation="46" />
        <Ladder price="47" representation="47" />
        <Ladder price="48" representation="48" />
        <Ladder price="49" representation="49" />
        <Ladder price="50" representation="50" />
        <Ladder price="52" representation="52" />
        <Ladder price="54" representation="54" />
        <Ladder price="56" representation="56" />
        <Ladder price="58" representation="58" />
        <Ladder price="60" representation="60" />
        <Ladder price="62" representation="62" />
        <Ladder price="64" representation="64" />
        <Ladder price="66" representation="66" />
        <Ladder price="68" representation="68" />
        <Ladder price="70" representation="70" />
        <Ladder price="72" representation="72" />
        <Ladder price="74" representation="74" />
        <Ladder price="76" representation="76" />
        <Ladder price="78" representation="78" />
        <Ladder price="80" representation="80" />
        <Ladder price="82" representation="82" />
        <Ladder price="84" representation="84" />
        <Ladder price="86" representation="86" />
        <Ladder price="88" representation="88" />
        <Ladder price="90" representation="90" />
        <Ladder price="92" representation="92" />
        <Ladder price="94" representation="94" />
        <Ladder price="96" representation="96" />
        <Ladder price="98" representation="98" />
        <Ladder price="100" representation="100" />
        <Ladder price="102" representation="102" />
        <Ladder price="104" representation="104" />
        <Ladder price="106" representation="106" />
        <Ladder price="108" representation="108" />
        <Ladder price="110" representation="110" />
        <Ladder price="112" representation="112" />
        <Ladder price="114" representation="114" />
        <Ladder price="116" representation="116" />
        <Ladder price="118" representation="118" />
        <Ladder price="120" representation="120" />
        <Ladder price="122" representation="122" />
        <Ladder price="124" representation="124" />
        <Ladder price="126" representation="126" />
        <Ladder price="128" representation="128" />
        <Ladder price="130" representation="130" />
        <Ladder price="132" representation="132" />
        <Ladder price="134" representation="134" />
        <Ladder price="136" representation="136" />
        <Ladder price="138" representation="138" />
        <Ladder price="140" representation="140" />
        <Ladder price="142" representation="142" />
        <Ladder price="144" representation="144" />
        <Ladder price="146" representation="146" />
        <Ladder price="148" representation="148" />
        <Ladder price="150" representation="150" />
        <Ladder price="152" representation="152" />
        <Ladder price="154" representation="154" />
        <Ladder price="156" representation="156" />
        <Ladder price="158" representation="158" />
        <Ladder price="160" representation="160" />
        <Ladder price="162" representation="162" />
        <Ladder price="164" representation="164" />
        <Ladder price="166" representation="166" />
        <Ladder price="168" representation="168" />
        <Ladder price="170" representation="170" />
        <Ladder price="172" representation="172" />
        <Ladder price="174" representation="174" />
        <Ladder price="176" representation="176" />
        <Ladder price="178" representation="178" />
        <Ladder price="180" representation="180" />
        <Ladder price="182" representation="182" />
        <Ladder price="184" representation="184" />
        <Ladder price="186" representation="186" />
        <Ladder price="188" representation="188" />
        <Ladder price="190" representation="190" />
        <Ladder price="192" representation="192" />
        <Ladder price="194" representation="194" />
        <Ladder price="196" representation="196" />
        <Ladder price="198" representation="198" />
        <Ladder price="200" representation="200" />
        <Ladder price="205" representation="205" />
        <Ladder price="210" representation="210" />
        <Ladder price="215" representation="215" />
        <Ladder price="220" representation="220" />
        <Ladder price="225" representation="225" />
        <Ladder price="230" representation="230" />
        <Ladder price="235" representation="235" />
        <Ladder price="240" representation="240" />
        <Ladder price="245" representation="245" />
        <Ladder price="250" representation="250" />
        <Ladder price="255" representation="255" />
        <Ladder price="260" representation="260" />
        <Ladder price="265" representation="265" />
        <Ladder price="270" representation="270" />
        <Ladder price="275" representation="275" />
        <Ladder price="280" representation="280" />
        <Ladder price="285" representation="285" />
        <Ladder price="290" representation="290" />
        <Ladder price="295" representation="295" />
        <Ladder price="300" representation="300" />
        <Ladder price="305" representation="305" />
        <Ladder price="310" representation="310" />
        <Ladder price="315" representation="315" />
        <Ladder price="320" representation="320" />
        <Ladder price="325" representation="325" />
        <Ladder price="330" representation="330" />
        <Ladder price="335" representation="335" />
        <Ladder price="340" representation="340" />
        <Ladder price="345" representation="345" />
        <Ladder price="350" representation="350" />
        <Ladder price="355" representation="355" />
        <Ladder price="360" representation="360" />
        <Ladder price="365" representation="365" />
        <Ladder price="370" representation="370" />
        <Ladder price="375" representation="375" />
        <Ladder price="380" representation="380" />
        <Ladder price="385" representation="385" />
        <Ladder price="390" representation="390" />
        <Ladder price="395" representation="395" />
        <Ladder price="400" representation="400" />
        <Ladder price="405" representation="405" />
        <Ladder price="410" representation="410" />
        <Ladder price="415" representation="415" />
        <Ladder price="420" representation="420" />
        <Ladder price="425" representation="425" />
        <Ladder price="430" representation="430" />
        <Ladder price="435" representation="435" />
        <Ladder price="440" representation="440" />
        <Ladder price="445" representation="445" />
        <Ladder price="450" representation="450" />
        <Ladder price="455" representation="455" />
        <Ladder price="460" representation="460" />
        <Ladder price="465" representation="465" />
        <Ladder price="470" representation="470" />
        <Ladder price="475" representation="475" />
        <Ladder price="480" representation="480" />
        <Ladder price="485" representation="485" />
        <Ladder price="490" representation="490" />
        <Ladder price="495" representation="495" />
        <Ladder price="500" representation="500" />
        <Ladder price="505" representation="505" />
        <Ladder price="510" representation="510" />
        <Ladder price="515" representation="515" />
        <Ladder price="520" representation="520" />
        <Ladder price="525" representation="525" />
        <Ladder price="530" representation="530" />
        <Ladder price="535" representation="535" />
        <Ladder price="540" representation="540" />
        <Ladder price="545" representation="545" />
        <Ladder price="550" representation="550" />
        <Ladder price="555" representation="555" />
        <Ladder price="560" representation="560" />
        <Ladder price="565" representation="565" />
        <Ladder price="570" representation="570" />
        <Ladder price="575" representation="575" />
        <Ladder price="580" representation="580" />
        <Ladder price="585" representation="585" />
        <Ladder price="590" representation="590" />
        <Ladder price="595" representation="595" />
        <Ladder price="600" representation="600" />
        <Ladder price="605" representation="605" />
        <Ladder price="610" representation="610" />
        <Ladder price="615" representation="615" />
        <Ladder price="620" representation="620" />
        <Ladder price="625" representation="625" />
        <Ladder price="630" representation="630" />
        <Ladder price="635" representation="635" />
        <Ladder price="640" representation="640" />
        <Ladder price="645" representation="645" />
        <Ladder price="650" representation="650" />
        <Ladder price="655" representation="655" />
        <Ladder price="660" representation="660" />
        <Ladder price="665" representation="665" />
        <Ladder price="670" representation="670" />
        <Ladder price="675" representation="675" />
        <Ladder price="680" representation="680" />
        <Ladder price="685" representation="685" />
        <Ladder price="690" representation="690" />
        <Ladder price="695" representation="695" />
        <Ladder price="700" representation="700" />
        <Ladder price="705" representation="705" />
        <Ladder price="710" representation="710" />
        <Ladder price="715" representation="715" />
        <Ladder price="720" representation="720" />
        <Ladder price="725" representation="725" />
        <Ladder price="730" representation="730" />
        <Ladder price="735" representation="735" />
        <Ladder price="740" representation="740" />
        <Ladder price="745" representation="745" />
        <Ladder price="750" representation="750" />
        <Ladder price="755" representation="755" />
        <Ladder price="760" representation="760" />
        <Ladder price="765" representation="765" />
        <Ladder price="770" representation="770" />
        <Ladder price="775" representation="775" />
        <Ladder price="780" representation="780" />
        <Ladder price="785" representation="785" />
        <Ladder price="790" representation="790" />
        <Ladder price="795" representation="795" />
        <Ladder price="800" representation="800" />
        <Ladder price="805" representation="805" />
        <Ladder price="810" representation="810" />
        <Ladder price="815" representation="815" />
        <Ladder price="820" representation="820" />
        <Ladder price="825" representation="825" />
        <Ladder price="830" representation="830" />
        <Ladder price="835" representation="835" />
        <Ladder price="840" representation="840" />
        <Ladder price="845" representation="845" />
        <Ladder price="850" representation="850" />
        <Ladder price="855" representation="855" />
        <Ladder price="860" representation="860" />
        <Ladder price="865" representation="865" />
        <Ladder price="870" representation="870" />
        <Ladder price="875" representation="875" />
        <Ladder price="880" representation="880" />
        <Ladder price="885" representation="885" />
        <Ladder price="890" representation="890" />
        <Ladder price="895" representation="895" />
        <Ladder price="900" representation="900" />
        <Ladder price="905" representation="905" />
        <Ladder price="910" representation="910" />
        <Ladder price="915" representation="915" />
        <Ladder price="920" representation="920" />
        <Ladder price="925" representation="925" />
        <Ladder price="930" representation="930" />
        <Ladder price="935" representation="935" />
        <Ladder price="940" representation="940" />
        <Ladder price="945" representation="945" />
        <Ladder price="950" representation="950" />
        <Ladder price="955" representation="955" />
        <Ladder price="960" representation="960" />
        <Ladder price="965" representation="965" />
        <Ladder price="970" representation="970" />
        <Ladder price="975" representation="975" />
        <Ladder price="980" representation="980" />
        <Ladder price="985" representation="985" />
        <Ladder price="990" representation="990" />
        <Ladder price="995" representation="995" />
        <Ladder price="1000" representation="1000" />
      </GetOddsLadderResult>
    </GetOddsLadderResponse>
  </soap:Body>
</soap:Envelope>`
