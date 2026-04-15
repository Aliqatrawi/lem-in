
# 🐜 Lem-in: Ant Farm Simulator

## 📌 Overview

**Lem-in** is a Go-based simulation project that models an ant colony. Given a map consisting of rooms and tunnels, the objective is to find the optimal way to guide ants from the `##start` room to the `##end` room in **as few turns as possible**.

This project is focused on:
- Pathfinding
- Optimized ant distribution
- Input validation and parsing
- Simulation and output formatting

---

## 📁 Project Structure

```
lem-in/
├── main.go                          # Entry point for the program
├── go.mod                           # Module definition
├── parser/                          # File reading and parsing logic
│   └── parser.go                                             
├── control/                         # Core simulation and pathfinding logic
│   ├── bestCombinPaths.go
│   ├── findPaths.go
│   ├── sendAnts.go               
│   └── validPaths.go     
├── types/                           # Shared data structures (e.g., Room)
│   └── types.go  
├── examples/                        # Sample example files to test the program
│   ├── badexample00.txt
│   ├── badexample01.txt
│   ├── example00.txt  
│   ├── example01.txt    
│   ├── example02.txt  
│   ├── example03.txt 
│   ├── example04.txt  
│   ├── example05.txt 
│   ├── example06.txt           
│   └── example07.txt                  
```

---

## 🧠 Objectives

- ✅ Read and validate a file containing ant farm structure
- ✅ Build an internal graph representation
- ✅ Find **all non-overlapping paths** from `##start` to `##end`
- ✅ Choose the best path combination minimizing turns
- ✅ Distribute ants across paths intelligently
- ✅ Simulate movements and print each turn in proper format

---

## 📥 Input Format

An input file includes:

- Number of ants (first line)
- Special rooms:
  ```
  ##start
  start_room x y
  ##end
  end_room x y
  ```

### ⚠ Rules:
- Room names **must not** start with `L` or `#`
- Coordinates must be integers
- No duplicate room names or coordinates
- No duplicated or self-links

---

## 📤 Output Format

Each move should follow:
```
L<ant_id>-<room_name>
```

Each line represents one **turn**. Example:
```
L1-A L2-B
L1-C L2-C
L1-end L2-end
```

---

## ▶️ How to Run

```bash
go run . examples/example00.txt
```

---

## ✅ Example

**Input:**
```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

**Output:**
```
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```

---

## 🚫 Error Handling

The program will print:

```
ERROR: invalid data format
```

If input is:
- Missing start/end
- Contains invalid room or link format
- Has no valid path from `##start` to `##end`

---

## 🧪 Testing

You can test with various maps inside the `examples/` folder or create your own.

---

## 🚀 Bonus Ideas

- Create a visualizer to animate ant movements (using coordinates)
- Add map validation with specific error messages
- Show statistics (e.g., total turns, path lengths)

---

## 🔧 Technologies

- Language: [Go](https://golang.org/)
- Standard Library Only (no third-party packages)# lem-in
