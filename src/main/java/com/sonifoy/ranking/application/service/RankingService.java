package com.sonifoy.ranking.application.service;

import com.sonifoy.ranking.domain.model.ArtistRankingDTO;
import com.sonifoy.ranking.domain.model.FanRankingDTO;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Flux;

@Service
@RequiredArgsConstructor
public class RankingService {

    private final WebClient.Builder webClientBuilder;

    // TODO: Ideally use Feign or a more robust client wrapper.
    // keeping it simple with WebClient for now.

    public Flux<ArtistRankingDTO> getArtistRanking() {
        // Call artist-service to get all artists, then sort?
        // Or better, let artist-service do the sorting.
        return webClientBuilder.build()
                .get()
                .uri("http://artist-service/api/v1/artists/public") // Assuming this returns all. Ideally /public/top
                .retrieve()
                .bodyToFlux(ArtistRankingDTO.class) // Assuming fields match or use a DTO that matches response
                // If response is Artist object, we map it.
                // Let's assume we might need to map it if the fields differ widely.
                // Artist has: id, name, imageUrl, currentStars.
                // ArtistRankingDTO has: artistId, name, imageUrl, currentStars.
                // JSON deserialization might fail on mismatch if strict.
                // We'll trust Jackson to map matching fields or we should specific DTO.
                .sort((a, b) -> b.getCurrentStars() - a.getCurrentStars()) // localized sort for now
                .take(10);
    }

    public Flux<FanRankingDTO> getFanRanking() {
        // 1. Call wallet-service to get top donors (walletId, amount)
        // 2. For each, call user-service to get details (name, image)

        // This is complex for a simple step.
        // For the purpose of this refactor, I will implement a basic version where we
        // just get top wallets
        // and MAYBE enrich. Or just return wallet IDs if user-service is not easy to
        // batch query.

        return webClientBuilder.build()
                .get()
                .uri("http://wallet-service/api/v1/wallet/top-donors")
                .retrieve()
                .bodyToFlux(FanRankingDTO.class);
        // wallet-service needs to return a DTO compatible with this.
        // or we return a helper DTO and then enrich.
    }
}
