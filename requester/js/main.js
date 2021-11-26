class Variables {
    variableBag

    constructor(s) {
        this.variableBag = JSON.parse(s)
    }

    set(key, value) {
        this.variableBag[key] = value
    }

    get(key) {
        return this.variableBag[key]
    }
}

class Client {
    session;
    environment;
    collection;

    constructor(session, environment, collection) {
        this.session = new Variables(session)
        this.environment = new Variables(environment)
        this.collection = new Variables(collection)

    }
}

class Response {
    body;
    headers;
    status;
    contentType;

    constructor(body, headers, status) {
        try {
            this.body = JSON.parse(body)
        } catch (error) {
            this.body = body
        }
        this.headers = new Headers(headers)
        this.status = JSON.parse(status)
    }
}

class Headers {
    #data

    constructor(headers) {
        this.#data = JSON.parse(headers)
    }

    valueOf(s) {
        return this.#data[s][0] ?? null
    }

    valuesOf(s) {
        return this.#data[s] ?? null
    }
}

var client = new Client(`{{session}}`, `{{environment}}`, `{{collection}}`)
var response = new Response(`{{body}}`, `{{headers}}`, `{{status}}`)

// {{Custom Code}}

