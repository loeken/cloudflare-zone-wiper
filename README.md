# Cloudflare DNS Record Deletion Tool

This is a command-line tool for deleting all DNS records for a specified domain name in Cloudflare using the Cloudflare API.

## Prerequisites

To use this tool, you need to have the following:

- Go installed on your system (`download here: https://golang.org/dl/`)
- A Cloudflare account with API access enabled
- A Cloudflare API token or API key and email address with the necessary permissions to delete DNS records (see Cloudflare API Tokens documentation for details)

## Usage

1. Clone or download this repository to your local machine:

git clone https://github.com/loeken/cloudflare-zone-wiper.git


2. Navigate to the directory where you downloaded the files:

cd cloudflare-zone-wiper

1. Install the necessary dependencies:

go mod download


1. Set the `API_TOKEN` and `DOMAIN` environment variables with your Cloudflare API token and the domain name for which you want to delete DNS records:

export API_TOKEN=YOUR_CLOUDFLARE_API_TOKEN
export DOMAIN=example.com

1. Run the tool:

go run main.go


The tool will use the Cloudflare API to retrieve the zone ID for the specified domain name, then delete all DNS records for that zone.

## Note

- This tool will delete **all** DNS records for the specified domain name in Cloudflare. Use with caution and make sure you have a backup of your DNS records before running this tool.
- This tool is provided as-is and without warranty. Use at your own risk.

## Contributing

If you have any issues or suggestions for improving this tool, please submit an issue or pull request in the GitHub repository.