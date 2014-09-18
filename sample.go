package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-mruby"
)

func main() {
	mrb := mruby.NewMrb()
	defer mrb.Close()

	file, err := os.Open("./Samplefile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	src := string(data[:])
	fmt.Printf("src is:\n%v\n", src)

	inFoobar := false
	parsedConfig := make(map[string]string)

	mrb_foobarBlock := func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
		inFoobar = true

		args := m.GetArgs()
		_, err := m.Yield(args[0])
		if err != nil {
			panic(err.Error())
		}

		inFoobar = false
		return mruby.Nil, nil
	}
	mrb_buz := func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
		if !inFoobar {
			return mruby.Nil, nil
		}
		args := m.GetArgs()
		parsedConfig["buz"] = args[0].String()

		return m.TrueValue(), nil
	}
	mrb_fizz := func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
		if !inFoobar {
			return mruby.Nil, nil
		}
		args := m.GetArgs()
		parsedConfig["fizz"] = args[0].String()

		return m.TrueValue(), nil
	}

	kernel := mrb.KernelModule()
	kernel.DefineMethod("foobar_block", mrb_foobarBlock, mruby.ArgsBlock())
	kernel.DefineMethod("buz", mrb_buz, mruby.ArgsReq(1))
	kernel.DefineMethod("fizz", mrb_fizz, mruby.ArgsReq(1))

	_, err = mrb.LoadString(src)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("parsed data is:\n%v\n", parsedConfig)
}
