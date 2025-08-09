# 🧠 Go Learning Journey by Mujammal Ahmed

Welcome to my Go learning journal!  
I'm learning the Go programming language from scratch and documenting everything I learn — day by day — in this repository.  
My goal is to build real-world projects with Go and help others learn from my journey.

---

## 📅 Progress Overview

| Day         | Topics Covered                                                                                                                    |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------- |
| [Day 1](#day-1) | Variables, Conditional Statements, Functions (`init`, IIFE), Scope (package, local, block)                                    |
| [Day 2](#day-2) | Parameters vs Arguments, First-order Functions, Higher-order Functions, Closures, Function Types, Middleware & Practical Patterns |

---

## 📚 Daily Logs

<details id="day-1">
<summary><strong>Day 1 — Variables, Conditional Statements, Functions (`init`, IIFE), Scope</strong></summary>

**Topics:**
- Variables (`var`, `:=`)
- Conditional Statements (`if`, `else`)
- Functions: standard, `init()`, and IIFE-style anonymous functions
- Scope: package-level, local, and block scope

**Code Sample:**  
See `day1/main.go` for the full code.

**Key Learnings:**
- Go requires explicit declarations — very clean but strict.
- `init()` runs automatically before `main()`.
- You can simulate Immediately Invoked Function Expressions (IIFE) using anonymous functions.

**Resources:**
- [Init Function](https://www.youtube.com/watch?v=UuWkHIyvwi0)
- [Fundamentals of Functions](https://medium.com/@danielabatibabatunde1/fundamentals-of-functions-in-golang-df4dd0c3072f)

</details>

---

<details id="day-2">
<summary><strong>Day 2 — Parameters vs Arguments, First-order & Higher-order Functions, Closures, Function Types</strong></summary>

1. **Parameters vs Arguments**  
   In Go, parameters are the named variables defined in a function’s declaration, while arguments are the actual values you pass when calling the function.  
   Go uses **pass-by-value** by default — meaning the function gets a copy of the argument, not the original.  
   However, when you pass **pointers** or **reference types** (slices, maps, channels, functions), the function can modify the original data because both point to the same underlying memory.  

   **Key takeaway:** parameters are the “placeholders,” arguments are the “real data.” Whether the function can change the caller’s data depends on whether you pass by value or by reference.

---

2. **First-order Functions**  
   A first-order function:
   - Takes only data types (e.g., `int`, `string`, `struct`) as inputs.
   - Returns only data types — **not** functions.  

   These are perfect for straightforward tasks like calculations, string processing, or struct initialization.  

   **Key takeaway:** if a function never accepts or returns another function, it’s first-order.

---

3. **Higher-order Functions**  
   A higher-order function is one that:
   - Takes one or more functions as parameters, **or**
   - Returns a function, **or**
   - Does both.

   Go supports higher-order functions because functions are **first-class citizens** — they can be stored in variables, passed to other functions, or returned.  

   Common patterns:
   - **Closures** — returned functions that “remember” variables from their outer scope.
   - **Middleware** — wrapping handlers with extra logic (logging, authentication).
   - **Factories** — generating functions with pre-set behavior.

   **Key takeaway:** higher-order functions power flexibility, reusability, and composition.

---

4. **Closures**  
   A closure is a function value that references variables from outside its body.  
   Those variables are **captured** and remembered even after the outer function has returned.  

   **Key takeaway:** closures combine behavior and remembered state — useful for counters, ID generators, and configuration-based function builders.

---

5. **Function Types & nil Functions**  
   In Go, functions have **explicit types** based on parameter and return signatures (e.g., `func(int) int`).  
   Function variables **default to nil** — so always check before calling if they might not be set.

---

**Code Samples:**  
See `day2/` for the full code.
</details>

---


## 🔖 Resources I Use Regularly

- [The Go Programming Language Docs](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Go Wiki](https://go.dev/wiki/)

---

