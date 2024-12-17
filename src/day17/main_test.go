package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	type args struct {
		instruction Instruction
		registers   *Registers
		pc          int
	}
	tcs := []struct {
		desc   string
		input  args
		result Registers
	}{
		{
			desc: "If register C contains 9, the program 2,6 would set register B to 1",
			input: args{
				instruction: Instruction{
					Opcode:  2,
					Operand: 6,
				},
				registers: &Registers{
					C: 9,
				},
				pc: 0,
			},
			result: Registers{
				C: 9,
				B: 1,
				A: 0,
			},
		},
		{
			desc: "If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.",
			input: args{
				instruction: Instruction{
					Opcode:  4,
					Operand: 0,
				},
				registers: &Registers{
					B: 2024,
					C: 43690,
				},
				pc: 0,
			},
			result: Registers{
				C: 43690,
				B: 44354,
				A: 0,
			},
		},
		{
			desc: "If register B contains 29, the program 1,7 would set register B to 26.",
			input: args{
				instruction: Instruction{
					Opcode:  1,
					Operand: 7,
				},
				registers: &Registers{
					B: 29,
				},
				pc: 0,
			},
			result: Registers{
				C: 0,
				B: 26,
				A: 0,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			_ = run(tc.input.instruction, tc.input.registers, &tc.input.pc)
			if diff := cmp.Diff(tc.result, *tc.input.registers); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetOutput(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result string
	}{
		{
			desc: "If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.",
			input: []string{
				"Register A: 10",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 5,0,5,1,5,4",
			},
			result: "0,1,2",
		},
		{
			desc: "If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.",
			input: []string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			result: "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			desc: "After the above program halts, its final output will be 4,6,3,5,6,3,5,2,1,0.",
			input: []string{
				"Register A: 729",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			result: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			desc: "Your puzzle answer must be 1,6,7,4,3,0,5,0,6.",
			input: []string{
				"Register A: 63687530",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,3,7,5,0,3,1,5,4,1,5,5,3,0",
			},
			result: "1,6,7,4,3,0,5,0,6",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getOutput(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestFindARegister(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "This program outputs a copy of itself if register A is instead initialized to 117440. (The original initial value of register A, 2024, is ignored.)",
			input: []string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,3,5,4,3,0",
			},
			result: 117440,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := findARegister(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
