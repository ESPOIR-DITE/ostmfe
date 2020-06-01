function initMap(listener) {
    var myLatlng = {lat: -34.192697, lng: 18.427731};

    var map = new google.maps.Map(document.getElementById('map'), {zoom: 15, center: myLatlng});

    // Create the initial InfoWindow.
    var infoWindow = new google.maps.InfoWindow(
        {content: 'Simon\'s Town Museum !', position: myLatlng});
    infoWindow.open(map);
}