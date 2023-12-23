MAX_RED = 12
MAX_GREEN = 13
MAX_BLUE = 14

def get_possible_games_sum(input):
    
    possible_games_id_list = []

    powers_list = [] 

    with open(input) as input:
        exceed = False

        for line in input:

            if exceed:
                exceed = False
            
            max_r = 0
            max_g = 0
            max_b = 0

            # get game id
            split_1 = line.split(":")
            game_id = int(split_1[0].split(" ")[1])

            # divide by sets
            sets = split_1[1].split(";")
            
            for game_set in sets:

                game_set = game_set.strip()
                colors = game_set.split(",")                

                #divide by colors
                for color in colors:

                    color = color.strip()

                    # get color count and color name
                    color_count = int(color.split(" ")[0])
                    color_name = color.split(" ")[1]

                    # parse and add counters
                    if color_name == "red" and color_count > MAX_RED:
                        exceed = True
                        
                    elif color_name == "green" and color_count > MAX_GREEN:
                        exceed = True

                    elif color_name == "blue" and color_count > MAX_BLUE:
                        exceed = True

                    elif not (color_name == "red" and color_count <= MAX_RED or color_name == "green" and color_count <= MAX_GREEN or color_name == "blue" and color_count <= MAX_BLUE):   
                        return -1
                    
                    if color_name == "red":
                        max_r = max(max_r, color_count)
                    elif color_name == "green":
                        max_g = max(max_g, color_count)
                    elif color_name == "blue":
                        max_b = max(max_b, color_count)
                    
            power = max_r * max_g * max_b
            powers_list.append(power)
                
            if not exceed:
                possible_games_id_list.append(game_id)

            

    print(sum(powers_list))
    print(sum(possible_games_id_list))

    return 1

def main():
    get_possible_games_sum("input.txt")

if __name__ == "__main__":
    main()