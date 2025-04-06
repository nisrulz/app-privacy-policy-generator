#!/usr/bin/env bash

# App Privacy Policy Generator: A simple web app to generate a generic
# privacy policy for your Android/iOS apps
#
# Copyright 2017-Present Nishant Srivastava
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

# Load nvm into a shell session.
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

# Check and display the current version of Node Version Manager (NVM)
echo "🔍 Checking NVM version..."
nvm -v
echo ""

# Display the currently active Node.js version managed by NVM
echo "🌐 Current Node.js version:"
node -v
echo ""

# Use NVM to install or switch to the latest stable Node.js version
echo "⬆️ Installing latest stable Node.js version..."
nvm install node
echo ""

# Upgrade npm to the latest global version for system-wide package management enhancements
echo "📦 Upgrading npm to the latest global version..."
npm upgrade -g npm
echo ""

# Update all project-specific packages to their latest compatible versions as specified in package.json
echo "✨ Updating project-specific packages to their latest versions..."
npm upgrade

echo ""
