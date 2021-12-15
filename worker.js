// deployed on https://2022.umut.ninja
const ORIGIN = 'newyeartree.herokuapp.com';

addEventListener('fetch', event => {    
    var req = event.request;
    
    const modifiedRequest = new Request("https://"+ORIGIN, {
        body: req.body,
        headers: req.headers,
        method: req.method,
        redirect: req.redirect
    })
    if (req.headers.get('user-agent')?.includes("curl")){
        event.respondWith(fetch(modifiedRequest));
    } else {
        event.respondWith(Response.redirect("https://twitter.com/umutbasalt", 301))
    }
})