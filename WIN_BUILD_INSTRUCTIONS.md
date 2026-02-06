# Windows Build Instructions

These instructions document how to reproduce the build for the `breadcrumb.lh.dotted.nodailycal.24h` planner on Windows.

## 1. Prerequisites

You need to have the following installed. We used `scoop` for package management.

```powershell
# Install Scoop if you haven't already
# (See https://scoop.sh/ for installation instructions if needed)

# Install Go and MiKTeX
scoop install go miktex
```

## 2. Configuration

Create a custom configuration file at `cfg/custom_layout.yaml` with your specific layout settings.

**File:** `cfg/custom_layout.yaml`

```yaml
layout:
  paper:
    width: 15.5cm
    height: 20.7cm

    margin:
      top: 0.8cm
      bottom: 0.6cm
      left: 1.2cm
      right: 0.3cm

    reversemargins: true
    marginparwidth: .8cm
    marginparsep: 3mm

  lengths:
    quarterlyspring: \textcolor{white}{.}

  numbers:
    dailytodos: 6
    dailynotes: 27
    dailydiarygoals: 10
    dailydiarygrateful: 2
    dailydiarybest: 2
    dailydiarylog: 29
    notesonpage: 37
```

## 3. Build Process

The build consists of two steps:

1. Running the Go program (`plannergen`) to generate LaTeX source files.
2. Running `xelatex` to compile the LaTeX files into a PDF.

Run the following commands in **PowerShell**.

### Step 3a: Generate LaTeX Source

This command runs the Go generator using the `cfg/base.yaml`, `cfg/template_breadcrumb.yaml`, and your `cfg/custom_layout.yaml`.

```powershell
# Set the environment variable for the year
$env:PLANNER_YEAR = 2026

# Run the generator
go run cmd/plannergen/plannergen.go --config "cfg/base.yaml,cfg/template_breadcrumb.yaml,cfg/custom_layout.yaml"
```

*Note: If you run into "command not found" errors, you may need to ensure your `Path` includes the shim directories:*

```powershell
$env:Path = "C:\Users\$env:USERNAME\scoop\shims;C:\Users\$env:USERNAME\scoop\apps\miktex\current\texmfs\install\miktex\bin\x64;" + $env:Path
```

### Step 3b: Compile PDF

This command uses `xelatex` to compile the generated `.tex` files in the `out/` directory.

**Important:** You must enable automatic package installation for MiKTeX first, otherwise the build will pause waiting for input (which you won't see in a background process).

```powershell
# Enable auto-install for MiKTeX packages
initexmf --set-config-value=[MPM]AutoInstall=1

# Change to the 'out' directory
cd out

# Compile the document
# Note: We run this on 'custom_layout.tex' because that is the name derived from the last config file used.
xelatex -interaction=nonstopmode -synctex=1 custom_layout.tex
```

## 4. Output

The generated PDF will be located at:
`out/custom_layout.pdf`

You can rename it to your desired name:

```powershell
# (Assuming you are still in the 'out' folder)
copy custom_layout.pdf ..\release\breadcrumb.lh.dotted.nodailycal.24h.pdf
```
