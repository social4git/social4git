#!/bin/sh

# (1) Start by installing the command-line tool.

go install github.com/social4git/social4git/social4git@latest

# (2) Make yourself two git repos, one public, one private
#
#    For example, I use:
#         https://github.com/petar/petar.social4git.public.git
#         https://github.com/petar/petar.social4git.private.git
#    I'll use these in the examples below.
#
#    If people ask you what's your public social4git handle, its the URL to your public repo. In my case:
#         https://github.com/petar/petar.social4git.public.git
#    Needless to say, you can get creative and get yourself some cool DNS name for your handle, e.g.
#         starstarer.eth
#
#    Get yourself an access token (HTTP password) for these repos from your git provider.
#    Let's call it ACCESS_TOKEN.

# (3) Create a config file for your social identity

cat <<EOF >> ~/.social4git/config.json
{
     "handle": "https://github.com/petar/petar.social4git.public.git",
     "public_url": "https://github.com/petar/petar.social4git.public.git",
     "private_url": "https://github.com/petar/petar.social4git.private.git",
     "public_auth": { "access_token": "ACCESS_TOKEN" },
     "private_auth": { "access_token": "ACCESS_TOKEN" },
     "var_dir": "/Users/petar/.social4git"
}
EOF

# (4) You are ready to make your first post
#    There are three different ways you can make a post:

# Here's a quick way to make a short post:
social4git post -m 'Howdy friends! Welcome to non-federated decentralized social media.'
# This will output a link to the post you just created, it will look something like:
#    social4git-https://github.com/petar/petar.social4git.public.git?post=20230324175744_444563cf7e5d32c6eef15863c1dd2352fbbbffbf9a085c08f38b70bddbd8f611_fd01baf45c686d35981fe3ea31c9197a58d02c047f455b778b8311b93a851df0
#    You can give this link to anyone. They can retrieve the post using `social4git fetch` (shown below).

# Here's another way to make a longer post:
social4git post <<EOF
We are just getting started, so we still maintain presence on Twitter @social4git and @maymounkov
EOF

# And another:
echo 'If you want to `retweet` a post, just post its link into a post. (Awkward phrasing, I know.)' > post3.txt
social4git post -f post3.txt

# (5) View your own posts:

social4git show --my

# Explore the command-line flags. There are many options when it comes to viewing:

social4git show -h

# (6) Follow someone:

social4git follow --handle https://github.com/petar/petar.social4git.public.git

# (7) See who you're following right now:

social4git following

# (8) Unfollow someone

social4git unfollow --handle https://github.com/petar/petar.social4git.public.git

# (9) Fetch all posts from users you are following:

social4git sync

# This is a clever command. It will fetch the posts you're following and cache them both in your private repo (so you always have access to what you've once seen) and in a local cache on your machine for performance.

# (10) Now you can read the posts from folks you follow:

social4git show

# (11) You can always fetch the contents of a post, given a link to it:

social4git fetch --link social4git-https://github.com/petar/petar.social4git.public.git?post=20230324175744_444563cf7e5d32c6eef15863c1dd2352fbbbffbf9a085c08f38b70bddbd8f611_fd01baf45c686d35981fe3ea31c9197a58d02c047f455b778b8311b93a851df0

# (12) How do you `retweet`?
#    For now, just make a post containing a link to the post you want to `retweet`.

social4git post -m social4git-https://github.com/petar/petar.social4git.public.git?post=20230324175744_444563cf7e5d32c6eef15863c1dd2352fbbbffbf9a085c08f38b70bddbd8f611_fd01baf45c686d35981fe3ea31c9197a58d02c047f455b778b8311b93a851df0

# We'll add more convenient functionalities as we go.

# If you are wondering who to follow, we'll keep a list of users who want to be followed on:
#    https://github.com/social4git/social4git

# Thanks for being curious!
