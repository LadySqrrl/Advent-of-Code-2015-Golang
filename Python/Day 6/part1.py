rows, cols = (1000, 1000)
lights = [[False for i in range(cols)] for j in range(rows)]
coord1 = ""
coord2 = ""
lines = []

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
                        lights[i][j] = True
    elif "off" in line:
        for i in range(len(lights)):
            for j in range(len(lights[i])):
                if (i >= start1) & (i <= end1):
                    if (j >= start2) & (j <= end2):
                        lights[i][j] = False
    elif "toggle" in line:
        for i in range(len(lights)):
            for j in range(len(lights[i])):
                if (i >= start1) & (i <= end1):
                    if (j >= start2) & (j <= end2):
                        if lights[i][j] == True:
                            lights[i][j] = False
                        else:
                            lights[i][j] = True
    else:
        print("error, no action identified")

count = 0

for i in range(len(lights)):
            for j in range(len(lights[i])):
                if lights[i][j] == True:
                    count = count + 1

print(count)
                   