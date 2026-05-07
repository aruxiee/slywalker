
## 🎀 slywalker: Stealth Module Resolution & EDR Evasion

This project demonstrates difference between **Loud API Interaction** and **Stealthy Memory Parsing**. While standard scripts rely on the Windows API to find system libraries, slywalker utilizes a custom PEB parser to locate `kernel32.dll` without leaving a footprint in the Import Address Table (IAT) or triggering EDR hooks.

⚠️ **Please Note:** This project is strictly for **Educational and Authorized Penetration Testing**. I am not responsible for any of the shenanigans you guys pull.

---

## 📂 Project Overview

Slywalker consists of two different scripts that achieve the same goal — locating the memory address of `kernel32.dll`.

### 1. `slywalker.go`
*   **Mechanism**: Uses inline assembly to access the `GS` register and manually walks the PEB linked lists.
*   **Visibility**: Zero. It does not use any Windows APIs and leaves the IAT clean.
*   **Advantage**: Bypasses user-mode hooks and static scanners.

### 2. `flagged.go`
*   **Mechanism**: Uses `syscall.NewLazyDLL` and `GetModuleHandleW`, as a lot of exploit writers do.
*   **Visibility**: High. It explicitly requests resources from the Windows Loader.
*   **Disadvantage**: Easily intercepted by EDRs and flagged by behavioral analysis.

---

## 🚀 Execution

### Prerequisites
*   Windows 10/11 (x64)
*   Go Compiler ([get here](https://golang.org/dl/))

It is absolutely fine if you do not want to build binaries of them both on your system. I have included both the binaries in their respective folders alongside the Go scripts. 

- To run slywalker as a Go script,

```bash
go run .
```
- To run flagged as a Go script,

```bash
go run flagged.go
```
### Build slywalker

If you want to build/execute binaries, ensure you have `peb_amd64.s` in the same directory as your `slywalker.go`.
```powershell
go build -o slywalker.exe .
.\slywalker.exe
```

### Build flagged
```powershell
go build -o flagged.exe .\flagged.go
.\flagged.exe
```

**Note**: To see the impact, run `dumpbin /imports slywalker.exe` and observe the lack of system library imports compared to `flagged`.

---

## 噫 Security Implications & Impact

### **Implementing "flagged"**
Using standard APIs like `GetModuleHandle` or `LoadLibrary` creates telemetry.
- **Static Analysis**: Scanners see the function names in the binary and flag it.
- **Dynamic Analysis**: EDRs place hooks on these functions. When called, EDR inspects *Intent* of the call and drops the process.

### **Implementing "slywalker"**
Walking the PEB is like an direct system interrogation.
- **Bypassing Hooks**: It never calls *hooked* functions. You read the data directly from the kernel structure.
- **Signature Evasion**: Your code looks like harmless memory math to a heuristic scanner.
- **Reliability**: PEB is a mechanical necessity for Windows. It cannot be removed or blocked without breaking the OS.

---

## 🕵️ Engagement Utilization

In a Red Team engagement or penetration test, slywalker acts as the foundation for local execution stealth.

*   **Custom Loaders**: Use the PEB parser to find `kernel32.dll` and `ntdll.dll`, then manually resolve function addresses (like `VirtualAlloc` or `WriteProcessMemory`) to build a completely independent loader.
*   **Reflective Execution**: By finding the base addresses of loaded modules silently, you can perform reflective DLL injection or module overloading, making your payload appear as a legitimate part of a signed system process.
*   **Anti-Debugging**: The PEB contains the `BeingDebugged` flag. This script can be extended to detect researchers and change its behavior to avoid analysis.

---

## ｧｩ MITRE ATT&CK Mapping

This project demonstrates several techniques used by advanced adversaries.

| ID | Technique | Description |
| :--- | :--- | :--- |
| **T1027** | **Obfuscated Files or Information** | By removing imports from the IAT, the binary hides its capabilities from static analysis. |
| **T1106** | **Native API** | Bypassing standard wrappers to interact with the OS at a lower level. |
| **T1622** | **Debugger Evasion** | Using the PEB to identify and evade analysis environments. |
| **T1055** | **Process Injection** | It's the foundational step for locating modules needed to inject code into processes. |

---

<p align="center">
  With ❤️ by <b>Aradhya</b>
</p>
