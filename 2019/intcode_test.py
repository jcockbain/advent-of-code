from intcode import IntCode


def tests():
    program = [3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8]
    for i in [7, 8, 9]:
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i == 8))

    program = [3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8]
    for i in [7, 8, 9]:
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i < 8))

    program = [3, 3, 1108, -1, 8, 3, 4, 3, 99]
    for i in [7, 8, 9]:
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i == 8))

    program = [3, 3, 1107, -1, 8, 3, 4, 3, 99]
    for i in [7, 8, 9]:
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i < 8))

    for i in [0, 1, -2]:
        program = [3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9]
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i != 0))

        program = [3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1]
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        assert (output == (i != 0))

    program = [3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125,
               20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99]
    a = {7: 999, 8: 1000, 9: 1001}
    for i in [7, 8, 9]:
        vm = IntCode(program, lambda: i)
        vm.input()
        output = vm.run()
        # print(output)
        assert output == a[i]
    print('tests passed')


if __name__ == '__main__':
    tests()
