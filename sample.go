package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-mruby"
)

var inFoobar bool
var parsedConfig map[string]string

func makeMrbMappingFunc(kernel *mruby.Class, keyName string) mruby.Func {
	mrb_newMethod := func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
		args := m.GetArgs()
		fmt.Printf("got: %s = %s\n", keyName, args[0].String())
		if !inFoobar {
			fmt.Printf("but exit\n")
			return mruby.Nil, nil
		}
		parsedConfig[keyName] = args[0].String()

		return m.TrueValue(), nil
	}

	kernel.DefineMethod(keyName, mrb_newMethod, mruby.ArgsReq(1))
	return mrb_newMethod
}

func main() {
	inFoobar := false
	parsedConfig := make(map[string]string)

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

	kernel := mrb.KernelModule()
	kernel.DefineMethod("foobar_block", mrb_foobarBlock, mruby.ArgsBlock())

	makeMrbMappingFunc(kernel, "buz")
	makeMrbMappingFunc(kernel, "fizz")
	makeMrbMappingFunc(kernel, "wheezee")

	_, err = mrb.LoadString(src)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("parsed data is:\n%v\n", parsedConfig)
}
