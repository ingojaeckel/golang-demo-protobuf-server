protobuf.load("messages.proto", function(err, root) {
    if (err) throw err;

    var SomeRequest = root.lookup("messages.SomeRequest");
    var SomeResponse = root.lookup("messages.SomeResponse");

    var message = SomeRequest.create({ param: 123 });
    var buffer = SomeRequest.encode(message).finish();

    $(function() {
      $("#request").click(function() {
        $.ajax({
          type: "POST",
          url: "/test",
          data: buffer,
          dataType: "binary",
          responseType: "arraybuffer",
          processData: false,
          success: function(arrayBuffer) {
            var bytes = new Uint8Array(arrayBuffer)
            var resp = SomeResponse.decode(bytes);
            console.log(resp)
          },
          error: function(result) {
            console.warn("error: " + result);
          }
        });
      });
    });
});
