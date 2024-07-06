#!/bin/bash

# Function to append direnv hook to a shell configuration file
append_direnv_hook() {
    local shell_config=$1
    local hook_command=$2

    if [ -f "$shell_config" ]; then
        # Check if the hook command already exists in the config file
        if grep -qF "$hook_command" "$shell_config"; then
            echo "Direnv hook already present in $shell_config"
        else
            # Append the hook command to the end of the config file
            echo "$hook_command" >> "$shell_config"
            echo "Direnv hook added to $shell_config"
        fi
    else
        echo "Shell configuration file $shell_config not found."
    fi
}

# Append direnv hook to bash configuration file if present
append_direnv_hook "$HOME/.bashrc" 'eval "$(direnv hook bash)"'

# Append direnv hook to zsh configuration file if present
append_direnv_hook "$HOME/.zshrc" 'eval "$(direnv hook zsh)"'
