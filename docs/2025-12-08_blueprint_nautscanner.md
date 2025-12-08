# Goal
The goal of this application is to build an online first, but open source tool that you can log into. So we would need to create a website and then the service needs to run somewhere on the server where the user can basically add his repos (like his GitHub repos). The one thing it does is scan all the repos and it creates a beautiful list of all the versions that they have (like all the application libraries - React, its Next.js version, etc.). Then you can click on it to get information about the latest features (we need to have access to all the feature lists of all the libraries out there, grab them and display the major cool features).

The reason for that is what you find underneath in review - this is a new vulnerability that just popped up last Saturday for Next.js and for React. We need to also have a scraper web scraper that scrapes down all the vulnerabilities from certain pages or be subscribed to the page or whatever, so that we know about this beforehand. Then, people all can have an auto-scanning. Basically, you understand if any of their libraries are prone by any of these vulnerabilities.

Obviously, this is an open source version, so we will add a dockerized system that somebody can download and run locally in the system (whatever they want, no support). The online version we're going to charge a little bit of money for, in order to let people scan all their repos fully automatically and generate reports. 
# Architecture

The application will consist of the following components:

1. **Frontend**: A React-based web application that will serve as the user interface for the tool. The frontend will be built using Next.js for server-side rendering and improved performance.

2. **Backend**: A Node.js server that will handle the business logic and interact with the database. The backend will be responsible for scanning repositories, fetching library information, and generating reports.

3. **Database**: A PostgreSQL database to store user data, repository information, and scan results. The database will be designed to handle large amounts of data efficiently.

4. **Scraper**: A Python-based web scraper that will fetch the latest features and vulnerabilities from various sources. The scraper will run periodically to ensure the data is up-to-date.

5. **Dockerized System**: A Docker container that will allow users to run the application locally. The container will include all the necessary dependencies and configurations to run the application without any issues.

6. **Payment Gateway**: An integration with a payment gateway to handle the subscription and payment processing for the online version of the application.


# Review Example
https://aws.amazon.com/blogs/security/china-nexus-cyber-threat-groups-rapidly-exploit-react2shell-vulnerability-cve-2025-55182/?inf_ver=2&inf_ctx=o99AYqaPG1gOD86iPPAnZNyIyyFTwfRVwgZUHdf_C4X5alUnDTFismMnco6_IWchyI0lG5TVTD9GYLzsUUT8bR_GPwdtPvMVihS3cBKdOu2xqPNJdS6RVUzRi25WEoPdlADxlAiLApyariKkr8QODVSYmOE104d5eFwWsuJ3c1LL950UZpNsB4uQjqBIR-RAzGvuXqIrGiNnMHKZXFfcjGbdMz3EmDPMVup5rQC8H-4hCyT_PXjX5DoZaLJIZ9020KArQq-v5qZ_MwITFwL6T8WOyMY2H9iucO3rvOTd9a0yfreLiFgXX1YgGX0veIUAm_ep89DDhM1tBaVWMT3miA%3D%3D