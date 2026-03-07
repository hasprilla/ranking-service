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
public class FanRankingDTO {
    private Long userId;
    private String name;
    private BigDecimal totalStars; // Total amount donated
    private String imageUrl;
}
