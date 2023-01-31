// Gets called via onclick event for #hamburger
function toggleNav(elem) {
    // Get class list of current element and toggle given class to be there or not
    elem.classList.toggle("is-active");

    // Get sibling element of current one and do the same as above
    elem.nextElementSibling.classList.toggle("is-active");
}