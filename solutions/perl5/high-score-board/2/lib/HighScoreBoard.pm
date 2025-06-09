package HighScoreBoard;

use v5.40;

our %Scores;

sub set_player_scores (%new_scores) {
    %Scores = (%Scores, %new_scores);
    return
}

sub get_player_score ($player) {
    return %Scores{$player}
}

sub increase_player_scores (%additional_scores) {
    foreach my $name ( keys %additional_scores ) {
        $Scores{$name} += %additional_scores{$name} // 0
    }
    return
}

sub sort_players_by_name {
    return sort keys %Scores
}

sub sort_players_by_score {
    return sort { $Scores{$b} <=> $Scores{$a} } keys %Scores
}

sub delete_player ($player) {
    delete $Scores{$player};
    return
}

1