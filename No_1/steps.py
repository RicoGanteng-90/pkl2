def howManyStep(discs):    
    return 2 ** discs - 1

disc = int(input("Enter the amount of disk: "))
step = howManyStep(disc)

print("the amount of step to move", disc, "disk is", step)