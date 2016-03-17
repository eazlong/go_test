package qsort

func quickSort( values [] int, start int, end int ) {

	temp := values[start];
	p := start;
	i,j := start, end
	for i <= j {
		for j >= p && values[j] >= temp {
			j--;
		}
		if j >= p {
			values[p] = values[j];
			p = j
		}
		for i<=p && values[i] <= temp {
			i++;
		}
		if i <= p {
			values[p] = values[i];
			p = i;
		}
	}
	values[p] = temp;
	if p-start > 1 {
		quickSort( values, start, p-1 );
	} else if end-p > 1 {
		quickSort( values, p+1, end );
	}
}

func Qsort( values []int ) {
	if len(values) <=1 {
		return;
	}
	quickSort( values, 0, len(values)-1 );
}
