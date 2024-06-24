/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoi.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/25 00:37:43 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Atoi(s string) int {
	var t int

	flag := true
	isNegative := false
	for _, c := range s {
		if flag {
			if c == '-' {
				isNegative = !isNegative
				continue
			} else if c == '+' {
				continue
			}
			flag = false
		}
		if c < '0' || '9' < c {
			return 0
		}
		t = t*10 + int(c) - '0'
	}
	if isNegative {
		t = -t
	}
	return t
}
