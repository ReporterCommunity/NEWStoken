package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"reflect"
	"strconv"

	"strings"

	"time"

	"fmt"

	"context"

	"github.com/DaveAppleton/etherUtils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

var blkNo *big.Int
var blkSet bool

func showBool(label string, f func(*bind.CallOpts) (bool, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}
	fmt.Println(label, " : ", res)
	log.Println(label, " : ", res)
}
func showUInt16(label string, f func(*bind.CallOpts) (uint16, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}
	fmt.Println(label, " : ", res)
	log.Println(label, " : ", res)
}
func showAddress(label string, f func(*bind.CallOpts) (common.Address, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}
	fmt.Println(label, " : ", res.Hex())
	log.Println(label, " : ", res.Hex())
}

func showBigEther(label string, f func(*bind.CallOpts) (*big.Int, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}

	fmt.Println(label, " : ", etherUtils.EtherToStr(res))
	log.Println(label, " : ", etherUtils.EtherToStr(res))
}

func showBigInt(label string, f func(*bind.CallOpts) (*big.Int, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}
	fmt.Println(label, " : ", res)
	log.Println(label, " : ", res)
}
func showDate(label string, f func(*bind.CallOpts) (*big.Int, error)) {
	res, err := f(nil)
	if err != nil {
		fmt.Println(label, " : ", err)
		log.Println(label, " : ", err)
		return
	}
	tm := time.Unix(res.Int64(), 0)
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		fmt.Println(label, " : ", res, tm, err)
		log.Println(label, " : ", res, tm, err)
		return
	}

	fmt.Println(label, " : ", res, tm.In(loc))
	log.Println(label, " : ", res, tm.In(loc))
}

var special func()

func ltxe(tx *types.Transaction, err error) {
	if err != nil {
		log.Printf("err : %v\n", err)
		fmt.Printf("err : %v\n", err)
		stopTheRun()
		return
	}
	log.Printf("TxHash %s  \n", tx.Hash().Hex())

	waitForMined(tx)
	if special != nil {
		special()
	}
}

type event struct {
	Name string
	Hash string
}

var eventArray = []event{
	event{Name: "Transfer", Hash: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"},
	event{Name: "Approval", Hash: "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"},
	event{Name: "Fees", Hash: "0x3ffd609b9b5a2a69ee3777fb7c2198bfcaf6cd5475d9c06bea186ac64afc0831"},
	event{Name: "Purchase", Hash: "0x5bc97d73357ac0d035d4b9268a69240988a5776b8a4fcced3dbc223960123f40"},
	event{Name: "Reduction", Hash: "0x1cb2dbf66c8795a5a1001c63a30fc3d4c207b969d6836f6f600cbd64aac9cf28"},
	event{Name: "purchase", Hash: "0x89991552bdd88d5235d19000be98e1812c0226db7d09fc0c72bba4ce0f4c9f21"},
	event{Name: "Reduction", Hash: "0x3a9c0a69270824c03aaa322e1895173d468382eabf49b766dc1d519735b7a84e"},
	event{Name: "reduction", Hash: "0x91759657c2dea7b0557513934a5f750d247fb06508f4002ecad6aab7ec15c568"},
	event{Name: "DebugTheBloodyThing", Hash: "0xe4f3dddffa806392251759cf30c032d60c89bdf5f9031eaa9a4aa3fc5ea32403"},
	event{Name: "TellMummy", Hash: "0x7b9997f79265015b0a8e912d7edf3c142c2444a18da8139c5fdad34a2316e603"},
	event{Name: "Allocation", Hash: "0x2ccf21bc8a43b499670fe41c33ca0f7b56c83863aca7c1494f0ede9068d2731a"}, //Allocation(uint256 amount, uint256 date)
	event{Name: "FeeOnAllocation", Hash: "0x2a0213f6a9134dfad8fe565f70d7764a32fe66ccfa431bd1f981d6a4db282680"},
	event{Name: "TimeHack", Hash: "0x8f0b1826dc44e92fc659210a47f2903e35755db23852a67e4b0ba444d760a7e3"},
	event{Name: "TokenMinted", Hash: "0xb9144c96c86541f6fa89c9f2f02495cccf4b08cd6643e26d34ee00aa586558a8"},
	event{Name: "TokenBurned", Hash: "0x1af5163f80e79b5e554f61e1d052084d3a3fe1166e42a265798c4e2ddce8ffa2"},
	event{Name: "PartComplete", Hash: "0xcadcbce497c428a85330c77795a87b336eeb93bc01259e2fea1cd998196f9350"},
	event{Name: "StillToGo", Hash: "0xc94f798321235fa17dd9603fb88bb2634abbb740fede709d211ad44ba8c63870"},
}

var eventFound map[string]bool
var possibleFailure bool

func waitForMined(tx *types.Transaction) {
	possibleFailure = false
	eventFound = make(map[string]bool)
	c, err := getClient()
	if err != nil {
		fmt.Println("getclient", err)
		return
	}
	c.Commit()
	blkNo = new(big.Int).Add(blkNo, big.NewInt(1))
	fmt.Println("block ", blkNo)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	for {

		rct, err := c.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(4 * time.Second)
				continue
			}
			fmt.Println("wait tx rct", err)
			return
		}
		fmt.Println("Status : ", rct.Status)
		warning := ""
		if rct.Status == 0 {
			warning = " ***** possible failure"
			possibleFailure = true
		}
		fmt.Println("mined ", tx.Gas(), rct.GasUsed, "logs: ", len(rct.Logs), warning)
		log.Println("mined ", tx.Gas(), rct.GasUsed, "logs: ", len(rct.Logs), warning)
		if len(rct.Logs) > 0 {
			fmt.Println("Log entries")
			for _, logRec := range rct.Logs {
				if len(logRec.Topics) > 0 {

					for i, topic := range logRec.Topics {
						if i == 0 {
							found := false
							for _, ev := range eventArray {
								eventFound[ev.Name] = false
								if topic == common.HexToHash(ev.Hash) {
									eventFound[ev.Name] = true
									fmt.Println(ev.Name)
									log.Println(ev.Name)
									found = true
									break
								}
							}
							if !found {
								fmt.Println("Topics")
								fmt.Println(topic.Hex())
								log.Println("Topics")
								log.Println(topic.Hex())
							}
						} else {
							fmt.Println(topic.Hex())
							log.Println(topic.Hex())
						}

					}
				}
				bytA := logRec.Data
				fmt.Println("Data")
				log.Println("Data")
				for {
					if len(bytA) == 0 {
						break
					}
					if len(bytA) < 32 {
						fmt.Printf("0x%x\n", bytA)
						log.Printf("0x%x\n", bytA)
						break
					}
					fmt.Printf("0x%x\n", bytA[:32])
					log.Printf("0x%x\n", bytA[:32])
					bytA = bytA[32:]
				}
			}
		}
		return
	}
}

type step struct {
	Comment   string
	Label     string
	Action    string
	Parameter string
	Second    string
	Gas       *big.Int
}

var script []step

func loadScript(file string) (err error) {
	scriptData, err := ioutil.ReadFile("scripts/" + file)
	if err != nil {
		log.Println("Cannot load page file ", err)
		return
	}
	if err = json.Unmarshal(scriptData, &script); err != nil {
		log.Fatal("LoadScript: JSON ", err)
	}
	return
}

var stopped bool
var skipDest string
var jumpToLabel string

func runScript() {
	variables = make(map[string]*big.Int)
	stopped = false
	skipDest = ""
	pass := 1
	for {
		for stepNo, r := range script {
			if len(jumpToLabel) != 0 {
				skipDest = jumpToLabel
				pass = 2
				jumpToLabel = ""
			}
			if len(skipDest) != 0 {
				if strings.Compare(skipDest, r.Label) == 0 {
					skipDest = ""
				} else {
					continue
				}
			}
			label := "          "
			if len(r.Label) > 0 {
				label = fmt.Sprintf("% 9s", r.Label+":")
			}
			fmt.Printf("%s % 8s % 8s % 8s\n", label, r.Action, r.Parameter, r.Second)
			log.Printf("%s % 8s % 8s % 8s\n", label, r.Action, r.Parameter, r.Second)
			if err := doStep(r); err != nil {
				fmt.Println("Stopped at line ", stepNo, " with error ", err)
				log.Println("Stopped at line ", stepNo, " with error ", err)
				return
			}
			if stopped {
				fmt.Println("Stopped at line ", stepNo)
				log.Println("Stopped at line ", stepNo)
				fmt.Printf("%s % 8s % 8s % 8s\n", label, r.Action, r.Parameter, r.Second)
				skript{}.Display("", "")
				return
			}
		}

		if pass == 1 {
			return
		}
		pass = 1

	}
}

func stopTheRun() {
	stopped = true
}

var timeOffset int64

func doStep(little step) (err error) {
	c, _ := getClient()

	c.Commit()
	if little.Gas == nil {
		little.Gas = new(big.Int).SetInt64(0)
	}

	apiR := reflect.ValueOf(skript{Gas: little.Gas})

	apiF := apiR.MethodByName(little.Action)
	if !apiF.IsValid() {
		return errors.New("Function not found : " + little.Action)
	}
	args := []reflect.Value{reflect.ValueOf(little.Parameter), reflect.ValueOf(little.Second)}
	apiF.Call(args)
	return nil
}

type skript struct {
	Gas *big.Int
}

func (s skript) BuySaleToken(who string, value string) {
	v, ok := etherUtils.StrToEther(value)
	if !ok {
		log.Println("invalid value to send : ", value)
		stopTheRun()
		return
	}

	s.Gas = big.NewInt(600000)

	log.Println(who, "=", userAddress(who).Hex())
	ltxe(sendEthereum(userKey(who), saleAddress, v, s.Gas))
}

func (skript) FundUser(who string, value string) {
	vInEther, ok := etherUtils.StrToEther(value)

	if !ok {
		log.Println("invalid value to send : ", value)
		stopTheRun()
		return
	}
	pc, err := getClient()
	if err != nil {
		log.Println(err)
		stopTheRun()
		return
	}

	vInEtherWithGas := new(big.Int).Add(vInEther, etherUtils.PointOneEther())
	//deadline := time.Now().Add(20 * time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), deadline)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	bal, err := pc.BalanceAt(ctx, userAddress(who), nil)
	if err != nil {
		log.Println(err)
		stopTheRun()
		return
	}
	diff := new(big.Int).Sub(vInEtherWithGas, bal)
	if diff.Sign() < 1 {
		log.Println(who, " already has ", bal, ", not topping up.")
		return
	}
	log.Println("adding ", diff, " to ", who, " (", userAddress(who).Hex(), ")")

	ltxe(sendEthereum(roleKey("banker"), userAddress(who), diff, new(big.Int)))
}

func makeTime(muliplier string, units string) (res *big.Int, err error) {
	mult, ok := new(big.Int).SetString(muliplier, 10)
	if !ok {
		err = errors.New("not a number : " + muliplier)
		log.Println("cannot understand ", muliplier, " as a number")
		return
	}
	ttime := new(big.Int)
	if strings.Compare(units, "week") == 0 {
		ttime.SetUint64(7 * 24 * 60 * 60)
	} else if strings.Compare(units, "day") == 0 {
		ttime.SetUint64(24 * 60 * 60)
	} else if strings.Compare(units, "hour") == 0 {
		ttime.SetUint64(60 * 60)
	} else if strings.Compare(units, "minute") == 0 {
		ttime.SetUint64(60)
	} else if strings.Compare(units, "second") == 0 {
		ttime.SetUint64(1)
	} else {
		log.Println(units, " is not valid time unit (week/day/hour/minute")
		// error
		err = errors.New("bad unit " + units)
	}
	res = ttime.Mul(ttime, mult)
	return
}

func (skript) WaitSaleNoFunding(string, string) {
	c, _ := getClient()
	end, err := reporterTokenSale.EndTimestamp(nil)
	if err != nil {
		log.Println("get end : ", err)
		stopTheRun()
		return
	}
	now, err := canary.TimeStamp(nil)
	if err != nil {
		fmt.Println(err)
	}
	timeOffset := new(big.Int).Sub(end, now).Int64()
	c.AdjustTime(time.Duration(timeOffset) * time.Second)
}

func (skript) WaitSaleFunding(string, string) {
	c, _ := getClient()
	start, err := reporterTokenSale.StartTimestamp(nil)
	if err != nil {
		log.Println("get start : ", err)
		stopTheRun()
		return
	}
	now, err := canary.TimeStamp(nil)
	if err != nil {
		fmt.Println(err)
	}
	timeOffset := new(big.Int).Sub(start, now).Int64()
	fmt.Println(" adjust by ", timeOffset)
	c.AdjustTime(time.Duration(timeOffset) * time.Second)
	c.Commit()
	now2, err := canary.TimeStamp(nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("from ", now, " by ", timeOffset, " to ", now2, "to hit ", start)

}

func (skript) Jump(dest string, dummy string) {
	skipDest = dest
}

// type balRec struct {
// 	Name       string
// 	Addr       string
// 	Ether      string
// 	HGT        string
// 	AllocShare string
// }

func balanceString(name string, who common.Address) (balStr string, rec []string, isZero bool) {

	c, err := getClient()
	if err != nil {
		balStr = err.Error()
		return
	}
	rec = append(rec, name)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	bal, err := c.BalanceAt(ctx, who, nil)
	if err != nil {
		balStr = err.Error()
		return
	}
	b := etherUtils.EtherToStr(bal)
	rec = append(rec, b)

	bbal, _ := reporterToken.BalanceOf(nil, who)
	bb := etherUtils.EtherToStr(bbal)
	rec = append(rec, bb)

	balStr = b + "," + bb

	zero := new(big.Int).SetUint64(0)
	isZero = (bal.Cmp(zero) == 0) && (bbal.Cmp(zero) == 0)

	return
}

func (s skript) ListBalances(csvf string, dummy2 string) {
	var recs [][]string

	files, err := ioutil.ReadDir("adminKeys/")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		b, rec, _ := balanceString(file.Name(), roleAddress(file.Name()))

		fmt.Println(file.Name(), " ", roleAddress(file.Name()).Hex(), b)
		log.Println(file.Name(), " ", roleAddress(file.Name()).Hex(), b)
		recs = append(recs, rec)
	}

	files, err = ioutil.ReadDir("userKeys/")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		b, rec, isz := balanceString(file.Name(), userAddress(file.Name()))
		if !isz {
			fmt.Println(file.Name(), " ", userAddress(file.Name()).Hex(), b)
			log.Println(file.Name(), " ", userAddress(file.Name()).Hex(), b)
			recs = append(recs, rec)
		}
	}
	fileName := viper.GetString("csv")
	if len(csvf) > 0 {
		fileName = csvf
	}

	fil, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("open ", fileName, err.Error())
		return
	}
	w := csv.NewWriter(fil)
	w.WriteAll(recs) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	return
}

func (skript) FinishSale(string, string) {

	ltxe(reporterTokenSale.FinishSale(ownerTx))

}

func (skript) Display(string, string) {
	fmt.Println("=====CROWD SALE==========================")
	fmt.Println("Address ", saleAddress.Hex())
	showBigEther("Ether Raised", reporterTokenSale.WeiRaised)
	showBigEther("NEWS sold in Crowdsale ", reporterTokenSale.TokenRaised)
	showBigEther("NEWS available in CS ", reporterTokenSale.TokensForSale)

	showDate("Start Date", reporterTokenSale.StartTimestamp)
	showDate("End Date", reporterTokenSale.EndTimestamp)
	fmt.Println("=====NEWS==========================")
	fmt.Println("NEWS address ", coinAddress.Hex())
	showBigEther("NEWS issued : ", reporterToken.TotalSupply)
	fmt.Println("========== Reserves & Stuff ================")
	showAddress("Crowdsale Multisig", reporterTokenSale.MultiSig)

}

func (skript) SendToken(tofrom string, val string) {
	c, _ := getClient()
	tfa := strings.Split(tofrom, ",")
	if len(tfa) < 2 {
		log.Println("not enough params (from,to)")
		fmt.Println("not enough params (from,to)")
		stopTheRun()
		return
	}
	to := tfa[1]
	from := tfa[0]
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	bal, _ := c.BalanceAt(ctx, userAddress(from), nil)
	fmt.Println(from, " has a balance of ", etherUtils.EtherToStr(bal))
	bal, _ = reporterToken.BalanceOf(nil, userAddress(from))
	fmt.Println(from, " has a token balance of ", etherUtils.EtherToStr(bal))
	value, ok := etherUtils.StrToEther(val)
	if !ok {
		log.Println(val, "is not a nice number")
		fmt.Println(val, "is not a nice number")
		stopTheRun()
		return
	}

	fmt.Println("send from ", from, " to ", to, " val = ", value)
	ltxe(reporterToken.Transfer(userTx(from), userAddress(to), value))
}

var variables map[string]*big.Int

func (skript) SetVar(name string, val string) {
	v, ok := new(big.Int).SetString(val, 10)
	if !ok {
		fmt.Println("Nasty Number ", val)
		log.Println("Nasty Number ", val)
		stopTheRun()
		return
	}
	variables[name] = v
}

func (skript) DecVarJumpNotZero(name string, label string) {
	v, ok := variables[name]
	if !ok {
		fmt.Println("Nasty Name ", name)
		log.Println("Nasty Name ", name)
		stopTheRun()
		return
	}
	if v.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Premature Zero ", name)
		log.Println("Premature Zero ", name)
		stopTheRun()
		return
	}
	v = new(big.Int).Sub(v, big.NewInt(1))
	fmt.Println(name, " now ", v)
	variables[name] = v
	if v.Cmp(big.NewInt(0)) != 0 {
		fmt.Println("Hit Zero, Jump ", label)
		log.Println("Hit Zero, Jump ", label)
		jumpToLabel = label
	}
}

func (s skript) BulkFund(number string, amount string) {
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
		stopTheRun()
		return
	}
	for j := 0; j < num; j++ {
		s.FundUser(fmt.Sprintf("T%04d", j+1), amount)
	}
}

func (s skript) BuySale800(string, string) {
	for j := 1; j < 801; j++ {
		amount := "10"
		if j%2 == 0 {
			amount = "9"
		}
		s.BuySaleToken(fmt.Sprintf("T%04d", j), amount)
	}
}

var ethAmounts = []string{"0", "0.1", "1", "9.99", "10", "10.1", "99.9", "100", "100.1", "500", "1000"}

func (s skript) BuySaleStaggered(start string, howMany string) {
	st, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println(err)
		stopTheRun()
		return
	}
	num, err := strconv.Atoi(howMany)
	if err != nil {
		fmt.Println(err)
		stopTheRun()
		return
	}

	for j := 0; j < num; j++ {
		amount := ethAmounts[j%len(ethAmounts)]
		account := fmt.Sprintf("T%04d", st+j)
		if (j/len(ethAmounts))%2 == 0 {
			s.Approve(account, "")
		}
		s.BuySaleToken(account, amount)
	}
}

// func (s skript) BuySaleToken(who string, value string) {
// 	v, ok := strToEther(value)
// 	if !ok {
// 		log.Println("invalid value to send : ", value)
// 		stopTheRun()
// 		return
// 	}

// 	s.Gas = big.NewInt(600000)

// 	log.Println(who, "=", userAddress(who).Hex())
// 	ltxe(sendEthereum(userKey(who), saleAddress, v, s.Gas))

// }

type specialRecord struct {
	Block     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	WeiRaised *big.Int
	Hardcap   *big.Int
	H20Sold   *big.Int
	Rate      *big.Int
	Timestamp *big.Int
}

var specialPreList []specialRecord
var specialList []specialRecord

func specialSaleFunction() {
	var sr specialRecord
	sr.Block = new(big.Int).Add(big.NewInt(0), blkNo)
	sr.StartTime, _ = reporterTokenSale.StartTimestamp(nil)
	sr.EndTime, _ = reporterTokenSale.EndTimestamp(nil)
	sr.WeiRaised, _ = reporterTokenSale.WeiRaised(nil)
	sr.H20Sold, _ = reporterToken.TotalSupply(nil)
	sr.Rate, _ = reporterTokenSale.Rate(nil)
	sr.Timestamp, _ = canary.TimeStamp(nil)
	specialList = append(specialList, sr)
}

func (skript) StartSaleSpecial(string, string) {
	special = specialSaleFunction
}

func writeSpecial(file string, list []specialRecord) {
	fil, err := os.OpenFile(file+".json", os.O_RDWR|os.O_CREATE, 0755)
	defer fil.Close()
	if err != nil {
		log.Println("open ", file, err.Error())
		return
	}
	w := json.NewEncoder(fil)
	if err = w.Encode(&specialList); err != nil {
		fmt.Println(err)
		log.Panicln(err)
		stopTheRun()
		return
	}
	lineArray := [][]string{{"Block", "EndTime", "Ether Raised", "HGT Sold", "HGT Max", "HGT/Ether", "Started?", "TimeStamp"}}
	for _, sp := range list {
		sl := []string{
			fmt.Sprint(sp.Block),
			fmt.Sprint(sp.StartTime),
			fmt.Sprint(sp.EndTime),
			fmt.Sprint(sp.WeiRaised),
			fmt.Sprint(sp.H20Sold),
			fmt.Sprint(sp.Rate),
			fmt.Sprint(sp.Timestamp),
		}
		lineArray = append(lineArray, sl)
	}
	cfil, err := os.OpenFile(file+".csv", os.O_RDWR|os.O_CREATE, 0755)
	defer cfil.Close()
	if err != nil {
		log.Println("open ", file+".csv", err.Error())
		return
	}
	cw := csv.NewWriter(cfil)
	cw.WriteAll(lineArray) // calls Flush internally

	if err := cw.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func (skript) WriteSaleSpecial(file string, dummy string) {
	writeSpecial(file, specialList)
}

func (skript) WritePresaleSpecial(file string, dummy string) {
	writeSpecial(file, specialPreList)
}

func (skript) WaitWeeks(numweeks string, dummy string) {
	nw, err := strconv.Atoi(numweeks)
	if err != nil {
		fmt.Println("Bad num weeks", numweeks, err)
		stopTheRun()
		return
	}
	oneWeek := time.Hour * 24 * 7
	client, err := getClient()
	client.AdjustTime(oneWeek * time.Duration(nw))
}

func (skript) Approve(whom string, dummy string) {
	key := userAddress(whom)
	ltxe(reporterTokenSale.AuthoriseAccount(ownerTx, key))
}
