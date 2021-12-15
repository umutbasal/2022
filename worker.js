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
	await LOGS.put(request.headers.get("cf-connecting-ip"), JSON.stringify({ headers: [...request.headers], url: request.url, method: request.method, date: Date.now() }))

	if (request.headers.get('user-agent')?.includes("curl")) {
		return fetch(request);
	} else {
		return Response.redirect("https://twitter.com/umutbasalt", 301);
	}
}