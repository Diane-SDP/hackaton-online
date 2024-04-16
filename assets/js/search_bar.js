document.getElementById('searchInput').addEventListener('input', function() {
    const searchValue = this.value.toLowerCase();
    const colisElements = document.querySelectorAll('.colis');

    colisElements.forEach(function(colis) {
        const uid = colis.getAttribute('data-uid').toLowerCase();
        if (uid.includes(searchValue)) {
            colis.style.display = 'block';
        } else {
            colis.style.display = 'none';
        }
    });
});