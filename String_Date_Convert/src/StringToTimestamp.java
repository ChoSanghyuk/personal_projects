import java.sql.Date;
import java.sql.Timestamp;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;

public class StringToTimestamp {
    public static void main(String[] args) throws ParseException {

        String sampleDate = "2023/09/05 23:22:11.123123";

        // 입력될 문자열의 포맷을 지정
        SimpleDateFormat parser = new SimpleDateFormat("yyyy/MM/dd HH:mm:ss.SSSD");
        // 결과로 도출될 문자열의 포맷을 지정
        SimpleDateFormat output = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss.SSSD");
        // parse : 문자열 => java.util.Date  format : java.util.Date => 문자열
        Timestamp date = Timestamp.valueOf(output.format(parser.parse(sampleDate)));

        System.out.println(date.toString()); // 2023-09-05

    }
}
