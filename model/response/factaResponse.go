package response

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Soapenc string   `xml:"soapenc,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                      string `xml:",chardata"`
		GetPartyFATCAInfoResponse struct {
			Text                    string `xml:",chardata"`
			P728                    string `xml:"p728,attr"`
			GetPartyFATCAInfoReturn struct {
				Text    string `xml:",chardata"`
				Control struct {
					Text          string `xml:",chardata"`
					Branch        string `xml:"branch"`
					Channel       string `xml:"channel"`
					RequestId     string `xml:"requestId"`
					RequesterName string `xml:"requesterName"`
					User          string `xml:"user"`
				} `xml:"control"`
				Status         string `xml:"status"`
				StatusDesc     string `xml:"statusDesc"`
				PartyFATCAInfo struct {
					Text       string `xml:",chardata"`
					AppEndDate struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"appEndDate"`
					AppFor               string `xml:"appFor"`
					LastUIDFATCA         string `xml:"lastUIDFATCA"`
					TaxRefNo             string `xml:"taxRefNo"`
					ChangeFATCAClassDate struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"changeFATCAClassDate"`
					CustomerId         string `xml:"customerId"`
					CustomerSource     string `xml:"customerSource"`
					EIN                string `xml:"EIN"`
					FATCAClass         string `xml:"FATCAClass"`
					FATCAStatus        string `xml:"FATCAStatus"`
					FTIN               string `xml:"FTIN"`
					GIIN               string `xml:"GIIN"`
					IRSFormExpiredDate struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"IRSFormExpiredDate"`
					IRSFormReceivedDate struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"IRSFormReceivedDate"`
					IRSFormSignedDate struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"IRSFormSignedDate"`
					IRSFormType       string `xml:"IRSFormType"`
					LatestFATCABranch string `xml:"latestFATCABranch"`
					Q1DocEvidence1    string `xml:"q1DocEvidence1"`
					Q1DocEvidenceFlag string `xml:"q1DocEvidenceFlag"`
					Q1Flag            string `xml:"q1Flag"`
					Q2DocEvidence1    string `xml:"q2DocEvidence1"`
					Q2DocEvidenceFlag string `xml:"q2DocEvidenceFlag"`
					Q2Flag            string `xml:"q2Flag"`
					Q3DocEvidence1    string `xml:"q3DocEvidence1"`
					Q3DocEvidence2    string `xml:"q3DocEvidence2"`
					Q3DocEvidenceFlag string `xml:"q3DocEvidenceFlag"`
					Q3Flag            string `xml:"q3Flag"`
					Q4DocEvidence1    string `xml:"q4DocEvidence1"`
					Q4DocEvidence2    string `xml:"q4DocEvidence2"`
					Q4DocEvidenceFlag string `xml:"q4DocEvidenceFlag"`
					Q4Flag            string `xml:"q4Flag"`
					Q5DocEvidence1    string `xml:"q5DocEvidence1"`
					Q5DocEvidence2    string `xml:"q5DocEvidence2"`
					Q5DocEvidenceFlag string `xml:"q5DocEvidenceFlag"`
					Q5Flag            string `xml:"q5Flag"`
					Q6DocEvidence1    string `xml:"q6DocEvidence1"`
					Q6DocEvidence2    string `xml:"q6DocEvidence2"`
					Q6DocEvidence3    string `xml:"q6DocEvidence3"`
					Q6DocEvidenceFlag string `xml:"q6DocEvidenceFlag"`
					Q6Flag            string `xml:"q6Flag"`
					Q7DocEvidence1    string `xml:"q7DocEvidence1"`
					Q7DocEvidence2    string `xml:"q7DocEvidence2"`
					Q7DocEvidenceFlag string `xml:"q7DocEvidenceFlag"`
					Q7Flag            string `xml:"q7Flag"`
					Q8DocEvidence1    string `xml:"q8DocEvidence1"`
					Q8DocEvidence2    string `xml:"q8DocEvidence2"`
					Q8DocEvidenceFlag string `xml:"q8DocEvidenceFlag"`
					Q8Flag            string `xml:"q8Flag"`
					Q9DocEvidence1    string `xml:"q9DocEvidence1"`
					Q9DocEvidence2    string `xml:"q9DocEvidence2"`
					Q9DocEvidenceFlag string `xml:"q9DocEvidenceFlag"`
					Q9Flag            string `xml:"q9Flag"`
					SSN               string `xml:"SSN"`
					B09Other          string `xml:"b09Other"`
					GIINSponsoring    string `xml:"GIINSponsoring"`
					NameSponsoring    string `xml:"nameSponsoring"`
					TaxResidency      string `xml:"taxResidency"`
					USCTP1            string `xml:"USCTP1"`
					USCTP1A1          string `xml:"USCTP1A1"`
					USCTP1A2          string `xml:"USCTP1A2"`
					USCTP1TIN         string `xml:"USCTP1TIN"`
					USCTP1CT          string `xml:"USCTP1CT"`
					USCTP2            string `xml:"USCTP2"`
					USCTP2A1          string `xml:"USCTP2A1"`
					USCTP2A2          string `xml:"USCTP2A2"`
					USCTP2TIN         string `xml:"USCTP2TIN"`
					USCTP2CT          string `xml:"USCTP2CT"`
					USCTP3            string `xml:"USCTP3"`
					USCTP3A1          string `xml:"USCTP3A1"`
					USCTP3A2          string `xml:"USCTP3A2"`
					USCTP3TIN         string `xml:"USCTP3TIN"`
					USCTP3CT          string `xml:"USCTP3CT"`
					USCTP4            string `xml:"USCTP4"`
					USCTP4A1          string `xml:"USCTP4A1"`
					USCTP4A2          string `xml:"USCTP4A2"`
					USCTP4TIN         string `xml:"USCTP4TIN"`
					USCTP4CT          string `xml:"USCTP4CT"`
					USCTP5            string `xml:"USCTP5"`
					USCTP5A1          string `xml:"USCTP5A1"`
					USCTP5A2          string `xml:"USCTP5A2"`
					USCTP5TIN         string `xml:"USCTP5TIN"`
					USCTP5CT          string `xml:"USCTP5CT"`
					USOwners          string `xml:"USOwners"`
					ActiveFlag        string `xml:"activeFlag"`
				} `xml:"partyFATCAInfo"`
			} `xml:"getPartyFATCAInfoReturn"`
		} `xml:"getPartyFATCAInfoResponse"`
	} `xml:"Body"`
}
