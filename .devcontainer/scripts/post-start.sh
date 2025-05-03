if command -v cursor &> /dev/null; then
    if [ -f ~/.zshrc ]; then
        echo "export GIT_EDITOR=\"cursor --wait\"" >> ~/.zshrc
    elif [ -f ~/.bashrc ]; then
        echo "export GIT_EDITOR=\"cursor --wait\"" >> ~/.bashrc
    else
        echo "export GIT_EDITOR=\"cursor --wait\"" >> ~/.profile
    fi
fi

if [ ! -f ./.devcontainer/.env.local ]; then
    touch ./.devcontainer/.env.local
fi

# Check if OPENAI_API_KEY is set in .env
if ! grep -q "OPENAI_API_KEY" ./.devcontainer/.env.local; then
    echo "OpenAI API key not found in ./.devcontainer/.env.local file."
    echo "📝 Please enter your OpenAI API key (https://platform.openai.com/api-keys):"
    read api_key
    
    echo "OPENAI_API_KEY=$api_key" >> ./.devcontainer/.env.local
    echo "🤖 The OpenAI API key has been saved to ./.devcontainer/.env.local file."
fi 


# Generate env file
rm .env && touch .env
cat ./.devcontainer/.env.dev >> .env
echo "" >> .env
cat ./.devcontainer/.env.local >> .env
