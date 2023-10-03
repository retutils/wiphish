# Wiphish: Advanced Phishing with Fake Access Point

`Wiphish` is a sophisticated red teaming tool that facilitates phishing attacks using a fake Access Point (AP), often referred to as the "Evil Twin" attack. The tool integrates a patched version of the renowned `Evilginx2` phishing framework and introduces a recursive DNS server. This ensures that target sites on phishing domains are only resolved once user credentials are captured, given that the user possesses a unique static IP.

**Table of Contents**
1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Safety and Ethics](#safety-and-ethics)
6. [Contribute](#contribute)
7. [License](#license)

## Features

1. **Evil Twin Attack**: Easily set up a fake AP to deceive unsuspecting users.
2. **Evilginx2 Patched Framework**: Incorporates the powerful features of Evilginx2, modified for enhanced phishing capabilities.
3. **Recursive DNS Server**: Ensures target phishing domains resolve only post-capture of credentials.
4. **Extended DHCP Lease**: Configured with a large DHCP lease time to accommodate users with unique static IPs.

## Prerequisites

- Hardware: Ensure you have compatible hardware for setting up a fake AP.
- Dependencies: List of required libraries/software (e.g., Evilginx2, DHCP server tools, etc.)

## Installation

```bash
# Clone the repository
git clone https://github.com/retutils/wiphish.git

# Navigate to the Wiphish directory
cd wiphish

# Install required dependencies (if any)
./install.sh
```

## Usage

Provide step-by-step usage instructions here, including any configuration files that need to be modified, command-line arguments, etc.

Example:
```bash
# Step 1: Setup your configuration
nano config.yml

# Step 2: Launch Wiphish
./wiphish.sh --target example.com
```

## Safety and Ethics

**Important**: Using `Wiphish` for unauthorized phishing attacks is illegal and unethical. Always obtain proper authorization before conducting any red team exercises. Use this tool responsibly and ethically.

## Contribute

Contributions are welcome! Please ensure your pull request adheres to the following guidelines:

- Clear, descriptive commit messages.
- Include relevant unit tests for any new features or bug fixes.
- Ensure any new code is documented.

## License

Specify your license here, for example, MIT, GPLv3, etc. Always ensure that your tool adheres to the licensing of the software/tools it builds upon.

---

Remember to replace placeholders like `yourusername` with appropriate values. Tailor the README to match the specific features, requirements, and steps of your tool. Ensure your tool adheres to the legal and ethical standards of the cybersecurity community.