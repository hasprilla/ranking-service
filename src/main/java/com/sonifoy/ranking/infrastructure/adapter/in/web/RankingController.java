package com.sonifoy.ranking.infrastructure.adapter.in.web;

import com.sonifoy.ranking.application.service.RankingService;
import com.sonifoy.ranking.domain.model.ArtistRankingDTO;
import com.sonifoy.ranking.domain.model.FanRankingDTO;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;

@RestController
@RequestMapping("/api/v1/ranking")
@RequiredArgsConstructor
public class RankingController {

    private final RankingService rankingService;

    @GetMapping("/artists")
    public Flux<ArtistRankingDTO> getArtistRanking() {
        return rankingService.getArtistRanking();
    }

    @GetMapping("/fans")
    public Flux<FanRankingDTO> getFanRanking() {
        return rankingService.getFanRanking();
    }
}
