function initWasm() {
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("payyans.wasm"), go.importObject).then((result) => {
    go.run(result.instance)
  });
}

function updateMap() {
  var mapUrl = $('#mapselect').val();
  mapUrl = './maps/' + mapUrl + '.map';
  $.ajax({
    url: mapUrl,
    success: function(data) {
      $('#map').val(data);
    },
    dataType: 'text'
  });
}

document.querySelector('#mapselect').addEventListener('change', function(e){
  updateMap()
});

document.querySelector('#convert-ascii').addEventListener('click', function(e){
  document.querySelector('#unicode-input').value = "Converting...";
  const fontMap = document.querySelector('#map').value;

  const result = AsciiToUnicodeByMapString(document.querySelector('#ascii-input').value, fontMap, "");
  
  document.querySelector('#unicode-input').value = result
});

initWasm()
updateMap()
