#!/bin/bash

# Funktion zum Ausführen oder Einbinden eines anderen Skripts
run_script() {
    local script=$1

    if [ -f "$script" ]; then
        # Ausführen des Skripts
        source "$script"
    else
        echo "Skript $script nicht gefunden."
    fi
}

# Aufrufen von setup_direnv.sh
run_script "./.devcontainer/setup_direnv.sh"

# Aufrufen von setup_go.sh
run_script "./.devcontainer/setup_go.sh"

echo "Setup abgeschlossen."
