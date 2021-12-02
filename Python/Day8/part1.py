def checkNum(assigned,num):
    try:
        return int(num)
    except:
        if num in assigned:
            return assigned[num]
        return None

def problem(probType):
    gates = {}
    input_file = open("input7.txt")
    for line in input_file:
        line = line.strip("\n")
        line = line.split(" ")
        if line[-1] in gates:
            gates[line[-1]].append(line[:-2])
        else:
            gates[line[-1]] = line[:-2]

    assigned = {}

    input_file.close()

    if probType == "b":
        assigned["b"] = 46065

    while len(assigned) < len(gates):
        if "a" in assigned:
            return assigned["a"]
        for k,v in gates.iteritems():
            item = v
            if k in assigned:
                continue
            if len(item) == 1:
                if not checkNum(assigned, item[0]) == None:
                    assigned[k] = checkNum(assigned, item[0])
            elif item[0] == "NOT":
                if not checkNum(assigned, item[1]) == None:
                    assigned[k] = ~ checkNum(assigned, item[1]) & 0xFFFF
            elif item[1] == "LSHIFT":
                if not checkNum(assigned, item[0]) == None:
                    assigned[k] = ((checkNum(assigned, item[0])) << int(item[2])) & 0xFFFF
            elif item[1] == "RSHIFT":
                if not checkNum(assigned, item[0]) == None:
                    assigned[k] = (checkNum(assigned, item[0]) >> int(item[2])) & 0xFFFF
            elif item[1] == "AND":
                if not checkNum(assigned, item[0]) == None and not checkNum(assigned,item[2]) == None:
                    assigned[k] = checkNum(assigned, item[0]) & checkNum(assigned, item[2]) & 0xFFFF
            elif item[1] == "OR":
                if not checkNum(assigned, item[0]) == None and not checkNum(assigned,item[2]) == None:
                    assigned[k] = checkNum(assigned, item[0]) | checkNum(assigned, item[2]) & 0xFFFF

    return assigned["a"]
                
print(problem("a"))
#print(problem("b"))