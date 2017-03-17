for (var i = 0; i < document.links.length; i++) {
    if (document.links[i].href == document.URL && !document.links[i].className.includes('skip-active')) {
        document.links[i].parentNode.className = 'active';
        if (document.links[i].parentNode.parentNode.className.includes('treeview-menu')) {
            document.links[i].parentNode.parentNode.parentNode.className = 'active';
        }
    }
}