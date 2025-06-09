package LanguageList;

use v5.40;

our @Languages;

sub add_language ($language) {
    push(@Languages, $language)
}

sub remove_language () {
    pop(@Languages)
}

sub first_language () {
    return $Languages[0]
}

sub last_language () {
    return $Languages[-1]
}

sub get_languages (@elements) {
    my @langs;
    foreach(@elements){
        push(@langs, $Languages[$_-1]);
    }
    return @langs
}

sub has_language ($language) {
    for(@Languages) {
        if ($_ eq $language) {
            return true
        }
    }
    return false
}

1;
