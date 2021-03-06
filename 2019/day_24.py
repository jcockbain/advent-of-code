import numpy as np


def get_nbs(point):
    return [(point[0] + dx, point[1] + dy) for dx, dy in zip([0, 1, 0, -1], [-1, 0, 1, 0])]


def iter(grid):
    new_grid = {}
    for point in grid:
        alive_nbs = 0
        nbs = get_nbs(point)
        for n in nbs:
            if n in grid and is_alive(grid[n]):
                alive_nbs += 1
        if is_alive(grid[point]):
            new_grid[point] = "#" if alive_nbs == 1 else "."
        else:
            new_grid[point] = "#" if 1 <= alive_nbs <= 2 else "."
    return new_grid


def iter_2(grid):
    pass


def print_world(grid):
    x, y = 0, 0
    H, W = get_dimensions_of_grid(grid)
    result = ""
    for point in grid:
        result += grid[point]
        x += 1
        if x == W + 1 and y != H:
            y += 1
            x = 0
            result += "\n"
    return result


def get_grid(world):
    grid = {}
    W, H = len(world[0]), len(world)
    grid = {(x, y): world[y][x] for y in range(H) for x in range(W)}
    return grid


def is_alive(symbol):
    return symbol == "#"


def get_dimensions_of_grid(grid):
    keys = list(grid.keys())
    keys.sort()
    H, W = keys[-1]
    return H, W


def p1(data):
    history = []
    grid = get_grid(data)
    while True:
        grid = iter(grid)
        if grid in history:
            break
        history.append(grid)
    return get_biodiversity_rating(grid)


def p2(data, minutes):
    grid = get_grid(data)
    return data


def get_biodiversity_rating(grid):
    power = 0
    rating = 0
    H, W = get_dimensions_of_grid(grid)
    for y in range(H+1):
        for x in range(W+1):
            if grid[(x, y)] == "#":
                rating += 2 ** power
            power += 1
    return rating


if __name__ == "__main__":
    with open("inputs/24.in", "r") as f:
        file = f.read()
    data = file.split("\n")
    print("part 1:  ", p1(data))
