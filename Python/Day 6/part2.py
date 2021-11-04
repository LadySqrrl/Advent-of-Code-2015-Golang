rows, cols = (1000, 1000)
lights = [[0 for i in range(cols)] for j in range(rows)]
coord1 = ""
coord2 = ""
lines = []
count = 0

with open('input.txt') as dir:
    lines = dir.readlines()

for line in lines:
    line = line.replace("\n", "")
    words = line.split(" ")
    for index in range(len(words)):
        if  (words[index] == "on") | (words[index] == "off") | (words[index] == "toggle"):
            coord1 = words[index + 1]
            
        if words[index] == 'through':
            coord2 = words[index + 1]
    
    #break the coords into two values
    start_coords = coord1.split(',')
    end_coords = coord2.split(',')
    start1 = int(start_coords[0])
    end1 = int(end_coords[0])
    start2 = int(start_coords[1])
    end2 =int(end_coords[1])
    
    # determine the correct action to take
    if "on" in line:
        for i in range(len(lights)):
            for j in range(len(lights[i])):
                if (i >= start1) & (i <= end1):
                    if (j >= start2) & (j <= end2):
                        lights[i][j] +=1
    elif "off" in line:
        for i in range(len(lights)):
            for j in range(len(lights[i])):
                if (i >= start1) & (i <= end1):
                    if (j >= start2) & (j <= end2):
                        if lights[i][j] > 0:
                            lights[i][j] -=1
    elif "toggle" in line:
        for i in range(len(lights)):
            for j in range(len(lights[i])):
                if (i >= start1) & (i <= end1):
                    if (j >= start2) & (j <= end2):
                        lights[i][j] +=2
    else:
        print("error, no action identified")

for i in range(len(lights)):
    for j in range(len(lights[i])):
        count += lights[i][j]

print(count)
                   