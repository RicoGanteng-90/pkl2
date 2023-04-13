def tower_of_hanoi(disk, first_rod, second_rod, third_rod):  
    if(disk == 1):  
        print('Move disk 1 from the', first_rod, 'rod to the', third_rod, 'rod')  
        return  
    
    tower_of_hanoi(disk - 1, first_rod, third_rod, second_rod)  
    print('Move disk', disk, 'from the', first_rod, 'rod to the', third_rod, 'rod')  
    tower_of_hanoi(disk - 1, second_rod, first_rod, third_rod)  
  
  
disk = int(input('Enter the number of disks: '))  

tower_of_hanoi(disk, 'first', 'second', 'third')
