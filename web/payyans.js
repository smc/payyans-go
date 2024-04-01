const fileUrl = path => {

}

const loadJS = function(url, implementationCode, location){
  var scriptTag = document.createElement("script");
  scriptTag.src = url;

  scriptTag.onload = implementationCode;
  scriptTag.onreadystatechange = implementationCode;

  location.appendChild(scriptTag);
};

const onScriptLoad = function(){
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("payyans.wasm"), go.importObject).then((result) => {
    go.run(result.instance)
  });
}

loadJS(fileUrl("wasm_exec.js"), onScriptLoad, document.body);
