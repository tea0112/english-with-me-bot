package statics

import "math/rand"

var MotivationalSentences = []string{
	"You can do it!",                          // Bạn có thể làm được!
	"Keep going, don't give up!",              // Tiếp tục cố gắng nhé!
	"I believe in you!",                       // Tôi tin bạn!
	"Stay strong.",                            // Cố gắng mạnh mẽ nhé.
	"Try one more time.",                      // Hãy thử lại lần nữa.
	"You are not alone.",                      // Bạn không cô đơn đâu.
	"Every day is a new chance.",              // Mỗi ngày là một cơ hội mới.
	"Small steps lead to big change.",         // Những bước nhỏ dẫn đến thay đổi lớn.
	"Your effort will pay off.",               // Nỗ lực của bạn sẽ được đền đáp.
	"Be proud of your progress.",              // Hãy tự hào về tiến bộ của bạn.
	"You did great!",                          // Bạn làm tốt lắm!
	"Tomorrow will be better.",                // Sẽ có ngày mai tốt hơn.
	"Never stop learning.",                    // Đừng bao giờ ngừng học hỏi.
	"Mistakes help you grow.",                 // Sai lầm giúp bạn trưởng thành.
	"You deserve success.",                    // Bạn xứng đáng thành công.
	"Believe in yourself.",                    // Hãy tin vào bản thân mình.
	"Your dreams matter.",                     // Ước mơ của bạn rất quan trọng.
	"Keep practicing every day.",              // Hãy luyện tập mỗi ngày.
	"You are making progress.",                // Bạn đang tiến bộ.
	"Don't give up!",                          // Đừng bỏ cuộc!
	"Go for it!",                              // Hãy làm như vậy đi!
	"It's okay to make mistakes.",             // Không sao khi phạm lỗi.
	"Failure teaches success.",                // Thất bại là mẹ thành công.
	"Each try makes you stronger.",            // Mỗi lần cố gắng khiến bạn mạnh mẽ hơn.
	"Focus on today.",                         // Tập trung vào hôm nay.
	"You can learn from every mistake.",       // Bạn có thể học từ mọi sai lầm.
	"Take a deep breath and try again.",       // Hãy hít sâu và thử lại.
	"Every expert was once a beginner.",       // Mọi chuyên gia từng là người mới bắt đầu.
	"You are stronger than you think.",        // Bạn mạnh mẽ hơn bạn nghĩ.
	"Practice makes perfect.",                 // Thực hành tạo nên sự hoàn hảo.
	"It's worth trying.",                      // Cũng đáng để thử đấy!
	"Hard work brings good results.",          // Làm việc chăm chỉ mang lại kết quả tốt.
	"Learning takes time.",                    // Học tập cần thời gian.
	"The first step is always hardest.",       // Bước đầu tiên luôn khó khăn nhất.
	"You grow when things are difficult.",     // Bạn phát triển khi gặp khó khăn.
	"Think positive thoughts.",                // Hãy suy nghĩ tích cực.
	"The journey matters more than the goal.", // Hành trình quan trọng hơn đích đến.
	"You are doing your best.",                // Bạn đang làm tốt nhất có thể.
	"One page a day makes a book.",            // Mỗi ngày một trang tạo nên cuốn sách.
	"Good things take time.",                  // Những điều tốt đẹp cần thời gian.
	"Be kind to yourself.",                    // Hãy tử tế với bản thân.
	"Today's effort is tomorrow's success.",   // Nỗ lực hôm nay là thành công ngày mai.
	"You learn something new every day.",      // Bạn học điều mới mỗi ngày.
	"Small wins add up.",                      // Những chiến thắng nhỏ cộng dồn lại.
	"Keep your eyes on your goals.",           // Giữ mắt vào mục tiêu của bạn.
	"You matter.",                             // Bạn rất quan trọng.
	"Start small, think big.",                 // Bắt đầu nhỏ, nghĩ lớn.
	"The best time to start is now.",          // Thời điểm tốt nhất để bắt đầu là bây giờ.
	"Your hard work is noticed.",              // Công sức của bạn được ghi nhận.
	"Step by step, you will get there.",       // Từng bước một, bạn
}

func GetRandomMotivationalSentence(motivations []string) string {
	return motivations[rand.Intn(len(MotivationalSentences))]
}
