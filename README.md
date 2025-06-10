# DolphinPeek 🐬

A lightweight desktop widget that displays system information with dolphin-themed humor and ASCII art animations.

## Features

### Current Implementation
- **System Information Display**: CPU/GPU temperature and utilization, RAM usage, disk space
- **Neofetch-style System Info**: OS details, architecture, and hardware overview
- **Weather Display**: Static weather information with ASCII art (API integration planned)
- **Dolphin Jokes & Facts**: Random rotation of dolphin-themed programming humor
- **Cycling Display**: Rotates through different information screens every 10 seconds
- **ASCII Art**: Weather symbols and dolphin animations in Nothing Phone aesthetic style

### Planned Features
- **Real-time System Monitoring**: Live CPU/GPU temps, RAM usage, disk space
- **Weather API Integration**: Live weather data via wttr.in (no authentication required)
- **Dolphin Animation**: ASCII art dolphin jumping out of ocean and looking at user
- **Windows Support**: Cross-platform compatibility
- **Configurable Timing**: Adjustable screen rotation intervals

## Architecture

### Technology Stack
- **Language**: Go (chosen for simplicity and cross-platform support over Rust)
- **GUI Framework**: Fyne (selected for ease of use and cross-platform compatibility)
- **Target Platforms**: Linux (primary), Windows (secondary)

### Project Structure
```
dolphinPeek/
├── main.go                 # Entry point and GUI setup
├── go.mod                  # Dependencies
├── internal/
│   ├── widget/
│   │   └── app.go         # Main widget logic and cycling
│   ├── screens/
│   │   ├── neofetch.go    # System info display
│   │   ├── weather.go     # Weather display
│   │   ├── animation.go   # Dolphin animation
│   │   └── jokes.go       # Random dolphin jokes/facts
│   ├── system/
│   │   └── info.go        # System data collection
│   └── config/
│       └── config.go      # Configuration management
└── README.md
```

### Design Decisions

**Language Choice**: Go over Rust
- Easier learning curve for new developers
- Simpler cross-platform development
- Good balance of performance and development speed
- Extensive standard library for system information

**GUI Framework**: Fyne
- Cross-platform compatibility out of the box
- Lightweight and resource-efficient
- Good support for real-time updates
- Simple API for basic widget functionality

**Architecture Pattern**: Separate implementations per platform
- Reduces complexity compared to conditional compilation
- Easier to maintain platform-specific code
- Cleaner separation of concerns

**Update Strategy**: Load-on-display rather than continuous real-time
- Reduces resource usage
- Updates information when screen is actually shown
- Balances freshness with efficiency

## Installation & Usage

### Prerequisites
- Go 1.21 or higher
- Linux (primary target) or Windows

### Setup
```bash
# Clone the repository
git clone <repository-url>
cd dolphinPeek

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

### Configuration
- **Window Size**: 400x300 pixels (compact widget size)
- **Cycle Interval**: 10 seconds (configurable in config/config.go)
- **Window Behavior**: Regular window (not always-on-top)
- **Resizing**: Disabled to maintain consistent widget appearance

## Screens

### 1. System Information (Neofetch-style)
- Hostname and username
- Operating system and architecture
- Go version
- CPU, GPU, Memory, and Disk information (placeholders for now)
- ASCII art mascot

### 2. Weather Display
- Current weather with ASCII art symbols
- Temperature, humidity, pressure
- Wind speed and visibility
- Dolphin-themed weather puns ("Dolphinitely sunny today!")
- Location: "Code Ocean"

### 3. Dolphin Jokes & Wisdom
- Random rotation of dolphin jokes and programming facts
- ASCII art dolphins
- Coding-related humor mixed with dolphin trivia
- "Dolphin Logic™" branding

### 4. Animation (Planned)
- Dolphin jumping out of ocean
- Dolphin popping head out to look at user
- Frame-based ASCII animation

## Development Roadmap

### Phase 1: Core Functionality ✅
- [x] Basic project structure
- [ ] GUI framework setup
- [ ] Screen cycling logic
- [ ] Static content displays
- [ ] Dolphin-themed humor integration

### Phase 2: Real System Data
- [ ] CPU temperature and utilization collection
- [ ] GPU temperature and utilization (NVIDIA support)
- [ ] RAM usage monitoring
- [ ] Disk space calculation
- [ ] Hostname and username detection

### Phase 3: Weather Integration
- [ ] wttr.in API integration
- [ ] Location detection
- [ ] Weather ASCII art mapping
- [ ] Error handling for network issues

### Phase 4: Animations
- [ ] Dolphin jumping animation frames
- [ ] Animation timing and transitions
- [ ] Smooth frame cycling

### Phase 5: Polish & Features
- [ ] Configuration file support
- [ ] Adjustable cycle timing
- [ ] Window positioning options
- [ ] Error handling and graceful degradation

### Phase 6: Cross-Platform
- [ ] Windows-specific system information collection
- [ ] Windows GUI testing and optimization
- [ ] Platform-specific installation guides

## Technical Challenges

### System Information Collection
- **Linux**: Reading from `/sys/class/thermal/` for CPU temps
- **GPU**: Integration with `nvidia-smi` for NVIDIA cards
- **Cross-platform**: Different APIs for Windows vs Linux
- **Permissions**: Handling cases where temp data isn't accessible

### Resource Efficiency
- **Memory Usage**: Minimizing RAM footprint for always-running widget
- **CPU Usage**: Efficient polling without excessive system calls
- **Update Frequency**: Balancing freshness with resource consumption

### Cross-Platform Compatibility
- **File System Differences**: Linux `/sys/` vs Windows registry/WMI
- **GUI Behavior**: Platform-specific window management
- **Dependencies**: Ensuring consistent behavior across platforms

## Contributing

This project is designed as a learning exercise for desktop development and cross-platform challenges. The focus is on:
- Clean, readable Go code
- Efficient resource usage
- Cross-platform compatibility
- Maintainable architecture

For feature requests or contributions open a github issue. The issue should state the wanted improvements in the form of user stories.  

---

*"Swimming through system stats, one thread at a time"* 🐬💻
