// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2018-04-17 09:45:54.737387 +0100 BST m=+0.113149917
package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/james-wilder/betdaq/soap"
)

const (
	ReadOnlyService = "http://api.betdaq.com/v2.0/ReadOnlyService.asmx"
	SecureService   = "https://api.betdaq.com/v2.0/Secure/SecureService.asmx"
)

type BetdaqClient struct {
	Username string
	Password string
}

func NewBetdaqClient(username string, password string) *BetdaqClient {
	return &BetdaqClient{
		Username: username,
		Password: password,
	}
}


func (c *BetdaqClient) CallCancelAllOrders (request CancelAllOrders) (*CancelAllOrdersResponse, error) {
	fmt.Println("CancelAllOrders")

	var (
		response CancelAllOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.CancelAllOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.CancelAllOrdersResult.ReturnStatus[0].Code,
			response.CancelAllOrdersResult.ReturnStatus[0].Description,
			response.CancelAllOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallCancelAllOrdersOnMarket (request CancelAllOrdersOnMarket) (*CancelAllOrdersOnMarketResponse, error) {
	fmt.Println("CancelAllOrdersOnMarket")

	var (
		response CancelAllOrdersOnMarketResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.CancelAllOrdersOnMarketResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.CancelAllOrdersOnMarketResult.ReturnStatus[0].Code,
			response.CancelAllOrdersOnMarketResult.ReturnStatus[0].Description,
			response.CancelAllOrdersOnMarketResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallCancelOrders (request CancelOrders) (*CancelOrdersResponse, error) {
	fmt.Println("CancelOrders")

	var (
		response CancelOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.CancelOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.CancelOrdersResult.ReturnStatus[0].Code,
			response.CancelOrdersResult.ReturnStatus[0].Description,
			response.CancelOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallChangeHeartbeatRegistration (request ChangeHeartbeatRegistration) (*ChangeHeartbeatRegistrationResponse, error) {
	fmt.Println("ChangeHeartbeatRegistration")

	var (
		response ChangeHeartbeatRegistrationResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ChangeHeartbeatRegistrationResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ChangeHeartbeatRegistrationResult.ReturnStatus[0].Code,
			response.ChangeHeartbeatRegistrationResult.ReturnStatus[0].Description,
			response.ChangeHeartbeatRegistrationResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallChangePassword (request ChangePassword) (*ChangePasswordResponse, error) {
	fmt.Println("ChangePassword")

	var (
		response ChangePasswordResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ChangePasswordResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ChangePasswordResult.ReturnStatus[0].Code,
			response.ChangePasswordResult.ReturnStatus[0].Description,
			response.ChangePasswordResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallDeregisterHeartbeat (request DeregisterHeartbeat) (*DeregisterHeartbeatResponse, error) {
	fmt.Println("DeregisterHeartbeat")

	var (
		response DeregisterHeartbeatResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.DeregisterHeartbeatResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.DeregisterHeartbeatResult.ReturnStatus[0].Code,
			response.DeregisterHeartbeatResult.ReturnStatus[0].Description,
			response.DeregisterHeartbeatResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetAccountBalances (request GetAccountBalances) (*GetAccountBalancesResponse, error) {
	fmt.Println("GetAccountBalances")

	var (
		response GetAccountBalancesResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.GetAccountBalancesResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetAccountBalancesResult.ReturnStatus[0].Code,
			response.GetAccountBalancesResult.ReturnStatus[0].Description,
			response.GetAccountBalancesResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetCurrentSelectionSequenceNumber (request GetCurrentSelectionSequenceNumber) (*GetCurrentSelectionSequenceNumberResponse, error) {
	fmt.Println("GetCurrentSelectionSequenceNumber")

	var (
		response GetCurrentSelectionSequenceNumberResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetCurrentSelectionSequenceNumberResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetCurrentSelectionSequenceNumberResult.ReturnStatus[0].Code,
			response.GetCurrentSelectionSequenceNumberResult.ReturnStatus[0].Description,
			response.GetCurrentSelectionSequenceNumberResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetEventSubTreeNoSelections (request GetEventSubTreeNoSelections) (*GetEventSubTreeNoSelectionsResponse, error) {
	fmt.Println("GetEventSubTreeNoSelections")

	var (
		response GetEventSubTreeNoSelectionsResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetEventSubTreeNoSelectionsResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetEventSubTreeNoSelectionsResult.ReturnStatus[0].Code,
			response.GetEventSubTreeNoSelectionsResult.ReturnStatus[0].Description,
			response.GetEventSubTreeNoSelectionsResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetEventSubTreeWithSelections (request GetEventSubTreeWithSelections) (*GetEventSubTreeWithSelectionsResponse, error) {
	fmt.Println("GetEventSubTreeWithSelections")

	var (
		response GetEventSubTreeWithSelectionsResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetEventSubTreeWithSelectionsResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetEventSubTreeWithSelectionsResult.ReturnStatus[0].Code,
			response.GetEventSubTreeWithSelectionsResult.ReturnStatus[0].Description,
			response.GetEventSubTreeWithSelectionsResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetMarketInformation (request GetMarketInformation) (*GetMarketInformationResponse, error) {
	fmt.Println("GetMarketInformation")

	var (
		response GetMarketInformationResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetMarketInformationResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetMarketInformationResult.ReturnStatus[0].Code,
			response.GetMarketInformationResult.ReturnStatus[0].Description,
			response.GetMarketInformationResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetOddsLadder (request GetOddsLadder) (*GetOddsLadderResponse, error) {
	fmt.Println("GetOddsLadder")

	var (
		response GetOddsLadderResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetOddsLadderResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetOddsLadderResult.ReturnStatus[0].Code,
			response.GetOddsLadderResult.ReturnStatus[0].Description,
			response.GetOddsLadderResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetOrderDetails (request GetOrderDetails) (*GetOrderDetailsResponse, error) {
	fmt.Println("GetOrderDetails")

	var (
		response GetOrderDetailsResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.GetOrderDetailsResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetOrderDetailsResult.ReturnStatus[0].Code,
			response.GetOrderDetailsResult.ReturnStatus[0].Description,
			response.GetOrderDetailsResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetPrices (request GetPrices) (*GetPricesResponse, error) {
	fmt.Println("GetPrices")

	var (
		response GetPricesResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetPricesResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetPricesResult.ReturnStatus[0].Code,
			response.GetPricesResult.ReturnStatus[0].Description,
			response.GetPricesResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallGetSPEnabledMarketsInformation (request GetSPEnabledMarketsInformation) (*GetSPEnabledMarketsInformationResponse, error) {
	fmt.Println("GetSPEnabledMarketsInformation")

	var (
		response GetSPEnabledMarketsInformationResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetSPEnabledMarketsInformationResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetSPEnabledMarketsInformationResult.ReturnStatus[0].Code,
			response.GetSPEnabledMarketsInformationResult.ReturnStatus[0].Description,
			response.GetSPEnabledMarketsInformationResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListAccountPostings (request ListAccountPostings) (*ListAccountPostingsResponse, error) {
	fmt.Println("ListAccountPostings")

	var (
		response ListAccountPostingsResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ListAccountPostingsResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListAccountPostingsResult.ReturnStatus[0].Code,
			response.ListAccountPostingsResult.ReturnStatus[0].Description,
			response.ListAccountPostingsResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListAccountPostingsById (request ListAccountPostingsById) (*ListAccountPostingsByIdResponse, error) {
	fmt.Println("ListAccountPostingsById")

	var (
		response ListAccountPostingsByIdResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ListAccountPostingsByIdResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListAccountPostingsByIdResult.ReturnStatus[0].Code,
			response.ListAccountPostingsByIdResult.ReturnStatus[0].Description,
			response.ListAccountPostingsByIdResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListBlacklistInformation (request ListBlacklistInformation) (*ListBlacklistInformationResponse, error) {
	fmt.Println("ListBlacklistInformation")

	var (
		response ListBlacklistInformationResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ListBlacklistInformationResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListBlacklistInformationResult.ReturnStatus[0].Code,
			response.ListBlacklistInformationResult.ReturnStatus[0].Description,
			response.ListBlacklistInformationResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListBootstrapOrders (request ListBootstrapOrders) (*ListBootstrapOrdersResponse, error) {
	fmt.Println("ListBootstrapOrders")

	var (
		response ListBootstrapOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ListBootstrapOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListBootstrapOrdersResult.ReturnStatus[0].Code,
			response.ListBootstrapOrdersResult.ReturnStatus[0].Description,
			response.ListBootstrapOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListMarketWithdrawalHistory (request ListMarketWithdrawalHistory) (*ListMarketWithdrawalHistoryResponse, error) {
	fmt.Println("ListMarketWithdrawalHistory")

	var (
		response ListMarketWithdrawalHistoryResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.ListMarketWithdrawalHistoryResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListMarketWithdrawalHistoryResult.ReturnStatus[0].Code,
			response.ListMarketWithdrawalHistoryResult.ReturnStatus[0].Description,
			response.ListMarketWithdrawalHistoryResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListOrdersChangedSince (request ListOrdersChangedSince) (*ListOrdersChangedSinceResponse, error) {
	fmt.Println("ListOrdersChangedSince")

	var (
		response ListOrdersChangedSinceResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.ListOrdersChangedSinceResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListOrdersChangedSinceResult.ReturnStatus[0].Code,
			response.ListOrdersChangedSinceResult.ReturnStatus[0].Description,
			response.ListOrdersChangedSinceResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListSelectionTrades (request ListSelectionTrades) (*ListSelectionTradesResponse, error) {
	fmt.Println("ListSelectionTrades")

	var (
		response ListSelectionTradesResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.ListSelectionTradesResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListSelectionTradesResult.ReturnStatus[0].Code,
			response.ListSelectionTradesResult.ReturnStatus[0].Description,
			response.ListSelectionTradesResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListSelectionsChangedSince (request ListSelectionsChangedSince) (*ListSelectionsChangedSinceResponse, error) {
	fmt.Println("ListSelectionsChangedSince")

	var (
		response ListSelectionsChangedSinceResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.ListSelectionsChangedSinceResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListSelectionsChangedSinceResult.ReturnStatus[0].Code,
			response.ListSelectionsChangedSinceResult.ReturnStatus[0].Description,
			response.ListSelectionsChangedSinceResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallListTopLevelEvents (request ListTopLevelEvents) (*ListTopLevelEventsResponse, error) {
	fmt.Println("ListTopLevelEvents")

	var (
		response ListTopLevelEventsResponse
	)

	err := c.doApiCall(request, &response, ReadOnlyService)
	if err != nil {
		return nil, err
	}

	if response.ListTopLevelEventsResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.ListTopLevelEventsResult.ReturnStatus[0].Code,
			response.ListTopLevelEventsResult.ReturnStatus[0].Description,
			response.ListTopLevelEventsResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallPlaceOrdersNoReceipt (request PlaceOrdersNoReceipt) (*PlaceOrdersNoReceiptResponse, error) {
	fmt.Println("PlaceOrdersNoReceipt")

	var (
		response PlaceOrdersNoReceiptResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.PlaceOrdersNoReceiptResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.PlaceOrdersNoReceiptResult.ReturnStatus[0].Code,
			response.PlaceOrdersNoReceiptResult.ReturnStatus[0].Description,
			response.PlaceOrdersNoReceiptResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallPlaceOrdersWithReceipt (request PlaceOrdersWithReceipt) (*PlaceOrdersWithReceiptResponse, error) {
	fmt.Println("PlaceOrdersWithReceipt")

	var (
		response PlaceOrdersWithReceiptResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.PlaceOrdersWithReceiptResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.PlaceOrdersWithReceiptResult.ReturnStatus[0].Code,
			response.PlaceOrdersWithReceiptResult.ReturnStatus[0].Description,
			response.PlaceOrdersWithReceiptResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallPulse (request Pulse) (*PulseResponse, error) {
	fmt.Println("Pulse")

	var (
		response PulseResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.PulseResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.PulseResult.ReturnStatus[0].Code,
			response.PulseResult.ReturnStatus[0].Description,
			response.PulseResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallRegisterHeartbeat (request RegisterHeartbeat) (*RegisterHeartbeatResponse, error) {
	fmt.Println("RegisterHeartbeat")

	var (
		response RegisterHeartbeatResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.RegisterHeartbeatResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.RegisterHeartbeatResult.ReturnStatus[0].Code,
			response.RegisterHeartbeatResult.ReturnStatus[0].Description,
			response.RegisterHeartbeatResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallSuspendAllOrders (request SuspendAllOrders) (*SuspendAllOrdersResponse, error) {
	fmt.Println("SuspendAllOrders")

	var (
		response SuspendAllOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.SuspendAllOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.SuspendAllOrdersResult.ReturnStatus[0].Code,
			response.SuspendAllOrdersResult.ReturnStatus[0].Description,
			response.SuspendAllOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallSuspendAllOrdersOnMarket (request SuspendAllOrdersOnMarket) (*SuspendAllOrdersOnMarketResponse, error) {
	fmt.Println("SuspendAllOrdersOnMarket")

	var (
		response SuspendAllOrdersOnMarketResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.SuspendAllOrdersOnMarketResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.SuspendAllOrdersOnMarketResult.ReturnStatus[0].Code,
			response.SuspendAllOrdersOnMarketResult.ReturnStatus[0].Description,
			response.SuspendAllOrdersOnMarketResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallSuspendFromTrading (request SuspendFromTrading) (*SuspendFromTradingResponse, error) {
	fmt.Println("SuspendFromTrading")

	var (
		response SuspendFromTradingResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.SuspendFromTradingResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.SuspendFromTradingResult.ReturnStatus[0].Code,
			response.SuspendFromTradingResult.ReturnStatus[0].Description,
			response.SuspendFromTradingResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallSuspendOrders (request SuspendOrders) (*SuspendOrdersResponse, error) {
	fmt.Println("SuspendOrders")

	var (
		response SuspendOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.SuspendOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.SuspendOrdersResult.ReturnStatus[0].Code,
			response.SuspendOrdersResult.ReturnStatus[0].Description,
			response.SuspendOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallUnsuspendFromTrading (request UnsuspendFromTrading) (*UnsuspendFromTradingResponse, error) {
	fmt.Println("UnsuspendFromTrading")

	var (
		response UnsuspendFromTradingResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.UnsuspendFromTradingResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.UnsuspendFromTradingResult.ReturnStatus[0].Code,
			response.UnsuspendFromTradingResult.ReturnStatus[0].Description,
			response.UnsuspendFromTradingResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallUnsuspendOrders (request UnsuspendOrders) (*UnsuspendOrdersResponse, error) {
	fmt.Println("UnsuspendOrders")

	var (
		response UnsuspendOrdersResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.UnsuspendOrdersResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.UnsuspendOrdersResult.ReturnStatus[0].Code,
			response.UnsuspendOrdersResult.ReturnStatus[0].Description,
			response.UnsuspendOrdersResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *BetdaqClient) CallUpdateOrdersNoReceipt (request UpdateOrdersNoReceipt) (*UpdateOrdersNoReceiptResponse, error) {
	fmt.Println("UpdateOrdersNoReceipt")

	var (
		response UpdateOrdersNoReceiptResponse
	)

	err := c.doApiCall(request, &response, SecureService)
	if err != nil {
		return nil, err
	}

	if response.UpdateOrdersNoReceiptResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.UpdateOrdersNoReceiptResult.ReturnStatus[0].Code,
			response.UpdateOrdersNoReceiptResult.ReturnStatus[0].Description,
			response.UpdateOrdersNoReceiptResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}


func (c *BetdaqClient) doApiCall(request, response interface{}, url string) error {
	soapRequest, err := soap.Encode(request, c.Username, c.Password)
	fmt.Println(string(soapRequest))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(soapRequest))

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("HTTP response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	err = soap.Decode(body, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

