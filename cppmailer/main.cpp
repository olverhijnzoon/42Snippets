#include <curl/curl.h>
#include <string>
#include <unistd.h>
#include <iostream>

size_t ReadCallback(void* ptr, size_t size, size_t nmemb, void* userp) {
    std::string* s = static_cast<std::string*>(userp);
    if (size * nmemb < 1 || s->empty())
        return 0;
    *(static_cast<char*>(ptr)) = s->front();
    s->erase(0, 1);
    return 1;
}

int main() {
    std::cout << "# 42Snippets" << std::endl;
    std::cout << "## Mail & Curl" << std::endl;

    /*
        The callback function is used as a custom data provider for libcurl's HTTP POST or FTP upload operations (in this specific case, for sending an email using SMTP protocol). It reads the contents of a std::string one character at a time, placing each character in the buffer provided by libcurl. This allows you to send a std::string as the body of the email using SMTP. The function also returns the number of units actually written (in this case, 1 for each character or 0 if there are no more characters), so libcurl knows when all data has been sent.
    */

    CURL *curl;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_ALL);

    while(true) {
        curl = curl_easy_init();
        if(curl) {
            struct curl_slist *recipients = NULL;
            std::string from = "from@mailhog.example";
            std::string to = "to@mailhog.example";
            std::string mailbody = "Subject: 42Snippets\r\n"
                                   "\r\n"
                                   "This is a test email.";

            // Set up connection parameters
            curl_easy_setopt(curl, CURLOPT_URL, "smtp://mailhog-smtp-service.default.svc.cluster.local:1025");
            curl_easy_setopt(curl, CURLOPT_MAIL_FROM, from.c_str());
            recipients = curl_slist_append(recipients, to.c_str());
            curl_easy_setopt(curl, CURLOPT_MAIL_RCPT, recipients);

            // Set up email body
            curl_easy_setopt(curl, CURLOPT_READDATA, &mailbody);
            curl_easy_setopt(curl, CURLOPT_READFUNCTION, ReadCallback);
            curl_easy_setopt(curl, CURLOPT_UPLOAD, 1L);

            // Send the email
            res = curl_easy_perform(curl);

            // Check for errors
            if(res != CURLE_OK) {
                fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
            }

            // Clean up
            curl_slist_free_all(recipients);
            curl_easy_cleanup(curl);
        }

        // Sleep for 30 seconds before the next loop
        sleep(30);
    }

    curl_global_cleanup();

    return (int)res;
}
