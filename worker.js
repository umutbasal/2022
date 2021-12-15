const ORIGIN = 'newyeartree.herokuapp.com';

addEventListener('fetch', event => {
	var req = event.request;

	const modifiedRequest = new Request("https://" + ORIGIN, {
		body: req.body,
		headers: req.headers,
		method: req.method,
		redirect: req.redirect
	});

	event.respondWith(handleRequest(modifiedRequest, event));
})

async function handleRequest(request, event) {
    let body;
	let reqclone = request.clone()
    if (request.body){
        let reader = reqclone.body.getReader();
        while (true) {
            const {done, value} = await reader.read();
            if (done) {
                break;
            }
            body = btoa(new TextDecoder().decode(value));
        }
    }
    await LOGS.put(request.headers.get("cf-connecting-ip")+"-"+Date.now(), JSON.stringify({ headers: [...request.headers], url: request.url, method: request.method, body: body, date: Date.now() }))

	if (request.headers.get('user-agent')?.includes("curl")) {
		return fetch(request);
	} else {
		return Response.redirect("https://twitter.com/umutbasalt", 301);
	}
}