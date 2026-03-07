package com.sonifoy.ranking.domain.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.math.BigDecimal;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ArtistRankingDTO {
    private Long artistId;
    private String name;
    private Integer currentStars;
    private String imageUrl;
}
